(**
 * Copyright (c) 2014, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the "flow" directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 *
 *)

open CommandUtils
open Sys_utils
open Utils
(* conversion *)

let dts_ext = ".d.ts"
let dts_ext_find_pattern = "*.d.ts"

type error_config = {
  color: Tty.color_mode;
  one_line: bool;
  show_all_errors: bool;
}

let call_succeeds try_function function_input =
  try
    try_function function_input;
    true
  with
  (* print failwith <msg> command's exception message *)
  | Failure msg -> let () = prerr_endline msg in
                   false
  | _ -> false

let convert_file error_config outpath file =
  let base = Filename.chop_suffix (Filename.basename file) dts_ext in
  let outpath = match outpath with
    | None -> Filename.dirname file | Some p -> p in
  let outfile = Filename.concat outpath base ^ ".js" in
  Printf.printf "converting %S -> %S\n%!" file outfile;
  let content = cat file in
  let ast, errors = Parser_dts.program_file ~fail:false content file in
  if errors = []
  then (
    let oc = open_out outfile in
    if
      let fmt = Format.formatter_of_out_channel oc in
      call_succeeds (Printer_dts.program fmt) ast
    then
      let () = Printf.printf "No errors!\n\n" in
      close_out oc;
      0, 1, 1
    else
      let () = Printf.printf "No errors!\n\n" in
      Printf.printf "Conversion was not successful!\n\n";
      close_out oc;
      Sys.remove outfile;
      0, 0, 1
  ) else (
    let n = List.length errors in
    Printf.printf "%d errors:\n" n;
    let flow_errors = List.map (fun e ->
      Errors_js.parse_error_to_flow_error e
    ) errors in
    Errors_js.print_error_summary
      ~one_line:error_config.one_line
      ~color:error_config.color
      (not error_config.show_all_errors)
      flow_errors;
    n, 0, 1
  )

  (* Printer_dts.program *)

let find_files_recursive path =
  Find.find_with_name [Path.make path] dts_ext_find_pattern

let find_files path =
  Array.fold_left (fun acc f ->
    if Filename.check_suffix f dts_ext
    then (Filename.concat path f) :: acc
    else acc
  ) [] (Sys.readdir path)

let sum f = List.fold_left (fun n i -> n + f i) 0

(* sum_triple adds triples (3-tuples) in a list obtained by applying
   function f on a list of inputs (similar to sum function which adds
   integers instead of triples) *)
let sum_triple f = List.fold_left (
  fun (x,y,z) i ->
    let a, b, c = f i
    in (a+x, b+y, c+z)) (0, 0, 0)

let convert_dir outpath path recurse error_config =
  let dts_files = if recurse
    then find_files_recursive path
    else find_files path in
  (* List.fold_left (convert_file outpath) dts_files *)
  sum_triple (convert_file error_config outpath) dts_files

let convert path recurse error_config outpath =
  let nerrs, successful_converts, total_files  =
    if Filename.check_suffix path dts_ext then (
      let outpath = match outpath with
        | None -> Some (Filename.dirname path)
        | _ -> outpath
      in
      convert_file error_config outpath path
    ) else (
      if recurse && outpath != None then
        failwith "output path not available when recursive";
      convert_dir outpath path recurse error_config
    ) in
  print_endlinef "Total Errors: %d\nTotal Files: %d\nSuccessful Conversions: %d"
    nerrs total_files successful_converts;
  (* exit code = number of unsuccessful coversions *)
  exit (total_files - successful_converts)

(* command wiring *)

let spec = {
  CommandSpec.
  name = "convert";
  doc = "";
  usage = Printf.sprintf
    "Usage: %s convert [DIR]\n\n\
      Convert *.d.ts in DIR if supplied, or current directory.\n\
      foo.d.ts is written to foo.js\n"
      CommandUtils.exe_name;
  args = CommandSpec.ArgSpec.(
    empty
    |> error_flags
    |> flag "--output" (optional string)
        ~doc:"Output path (not available when recursive)"
    |> flag "--r" no_arg
        ~doc:"Recurse into subdirectories"
    |> anon "dir" (optional string)
        ~doc:"Directory (default: current directory)"
  )
}

let main color one_line show_all_errors outpath recurse dir () =
  let path = match dir with
  | None -> "."
  | Some path -> path
  in
  let error_config = {
    color = parse_color_enum color;
    one_line;
    show_all_errors;
  } in
  if ! Sys.interactive
  then ()
  else
    SharedMem.(init default_config);
    convert path recurse error_config outpath

let command = CommandSpec.command spec main
