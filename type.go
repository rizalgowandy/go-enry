package slinguist

// CODE GENERATED AUTOMATICALLY WITH gopkg.in/src-d/enry.v1/internal/code-generator
// THIS FILE SHOULD NOT BE EDITED BY HAND
// Extracted from github/linguist commit: b6460f8ed6b249281ada099ca28bd8f1230b8892

var languagesType = map[string]Type{
	"1C Enterprise":    Programming,
	"ABAP":             Programming,
	"ABNF":             Data,
	"AGS Script":       Programming,
	"AMPL":             Programming,
	"ANTLR":            Programming,
	"API Blueprint":    Markup,
	"APL":              Programming,
	"ASN.1":            Data,
	"ASP":              Programming,
	"ATS":              Programming,
	"ActionScript":     Programming,
	"Ada":              Programming,
	"Agda":             Programming,
	"Alloy":            Programming,
	"Alpine Abuild":    Programming,
	"Ant Build System": Data,
	"ApacheConf":       Markup,
	"Apex":             Programming,
	"Apollo Guidance Computer": Programming,
	"AppleScript":              Programming,
	"Arc":                      Programming,
	"Arduino":                  Programming,
	"AsciiDoc":                 Prose,
	"AspectJ":                  Programming,
	"Assembly":                 Programming,
	"Augeas":                   Programming,
	"AutoHotkey":               Programming,
	"AutoIt":                   Programming,
	"Awk":                      Programming,
	"Batchfile":                Programming,
	"Befunge":                  Programming,
	"Bison":                    Programming,
	"BitBake":                  Programming,
	"Blade":                    Markup,
	"BlitzBasic":               Programming,
	"BlitzMax":                 Programming,
	"Bluespec":                 Programming,
	"Boo":                      Programming,
	"Brainfuck":                Programming,
	"Brightscript":             Programming,
	"Bro":                      Programming,
	"C":                        Programming,
	"C#":                       Programming,
	"C++":                      Programming,
	"C-ObjDump":                Data,
	"C2hs Haskell":             Programming,
	"CLIPS":                    Programming,
	"CMake":                    Programming,
	"COBOL":                    Programming,
	"COLLADA":                  Data,
	"CSON":                     Data,
	"CSS":                      Markup,
	"CSV":                      Data,
	"CWeb":                     Programming,
	"Cap'n Proto":              Programming,
	"CartoCSS":                 Programming,
	"Ceylon":                   Programming,
	"Chapel":                   Programming,
	"Charity":                  Programming,
	"ChucK":                    Programming,
	"Cirru":                    Programming,
	"Clarion":                  Programming,
	"Clean":                    Programming,
	"Click":                    Programming,
	"Clojure":                  Programming,
	"Closure Templates":        Markup,
	"CoffeeScript":             Programming,
	"ColdFusion":               Programming,
	"ColdFusion CFC":           Programming,
	"Common Lisp":              Programming,
	"Component Pascal":         Programming,
	"Cool":                     Programming,
	"Coq":                      Programming,
	"Cpp-ObjDump":              Data,
	"Creole":                   Prose,
	"Crystal":                  Programming,
	"Csound":                   Programming,
	"Csound Document":          Programming,
	"Csound Score":             Programming,
	"Cuda":                     Programming,
	"Cycript":                  Programming,
	"Cython":                   Programming,
	"D":                        Programming,
	"D-ObjDump":                Data,
	"DIGITAL Command Language": Programming,
	"DM":             Programming,
	"DNS Zone":       Data,
	"DTrace":         Programming,
	"Darcs Patch":    Data,
	"Dart":           Programming,
	"Diff":           Data,
	"Dockerfile":     Data,
	"Dogescript":     Programming,
	"Dylan":          Programming,
	"E":              Programming,
	"EBNF":           Data,
	"ECL":            Programming,
	"ECLiPSe":        Programming,
	"EJS":            Markup,
	"EQ":             Programming,
	"Eagle":          Markup,
	"Ecere Projects": Data,
	"Eiffel":         Programming,
	"Elixir":         Programming,
	"Elm":            Programming,
	"Emacs Lisp":     Programming,
	"EmberScript":    Programming,
	"Erlang":         Programming,
	"F#":             Programming,
	"FLUX":           Programming,
	"Factor":         Programming,
	"Fancy":          Programming,
	"Fantom":         Programming,
	"Filebench WML":  Programming,
	"Filterscript":   Programming,
	"Formatted":      Data,
	"Forth":          Programming,
	"Fortran":        Programming,
	"FreeMarker":     Programming,
	"Frege":          Programming,
	"G-code":         Data,
	"GAMS":           Programming,
	"GAP":            Programming,
	"GCC Machine Description": Programming,
	"GDB":      Programming,
	"GDScript": Programming,
	"GLSL":     Programming,
	"GN":       Data,
	"Game Maker Language":     Programming,
	"Genie":                   Programming,
	"Genshi":                  Programming,
	"Gentoo Ebuild":           Programming,
	"Gentoo Eclass":           Programming,
	"Gettext Catalog":         Prose,
	"Gherkin":                 Programming,
	"Glyph":                   Programming,
	"Gnuplot":                 Programming,
	"Go":                      Programming,
	"Golo":                    Programming,
	"Gosu":                    Programming,
	"Grace":                   Programming,
	"Gradle":                  Data,
	"Grammatical Framework":   Programming,
	"Graph Modeling Language": Data,
	"GraphQL":                 Data,
	"Graphviz (DOT)":          Data,
	"Groovy":                  Programming,
	"Groovy Server Pages":     Programming,
	"HCL":                      Programming,
	"HLSL":                     Programming,
	"HTML":                     Markup,
	"HTML+Django":              Markup,
	"HTML+ECR":                 Markup,
	"HTML+EEX":                 Markup,
	"HTML+ERB":                 Markup,
	"HTML+PHP":                 Markup,
	"HTTP":                     Data,
	"Hack":                     Programming,
	"Haml":                     Markup,
	"Handlebars":               Markup,
	"Harbour":                  Programming,
	"Haskell":                  Programming,
	"Haxe":                     Programming,
	"Hy":                       Programming,
	"HyPhy":                    Programming,
	"IDL":                      Programming,
	"IGOR Pro":                 Programming,
	"INI":                      Data,
	"IRC log":                  Data,
	"Idris":                    Programming,
	"Inform 7":                 Programming,
	"Inno Setup":               Programming,
	"Io":                       Programming,
	"Ioke":                     Programming,
	"Isabelle":                 Programming,
	"Isabelle ROOT":            Programming,
	"J":                        Programming,
	"JFlex":                    Programming,
	"JSON":                     Data,
	"JSON5":                    Data,
	"JSONLD":                   Data,
	"JSONiq":                   Programming,
	"JSX":                      Programming,
	"Jasmin":                   Programming,
	"Java":                     Programming,
	"Java Server Pages":        Programming,
	"JavaScript":               Programming,
	"Jison":                    Programming,
	"Jison Lex":                Programming,
	"Jolie":                    Programming,
	"Julia":                    Programming,
	"Jupyter Notebook":         Markup,
	"KRL":                      Programming,
	"KiCad":                    Programming,
	"Kit":                      Markup,
	"Kotlin":                   Programming,
	"LFE":                      Programming,
	"LLVM":                     Programming,
	"LOLCODE":                  Programming,
	"LSL":                      Programming,
	"LabVIEW":                  Programming,
	"Lasso":                    Programming,
	"Latte":                    Markup,
	"Lean":                     Programming,
	"Less":                     Markup,
	"Lex":                      Programming,
	"LilyPond":                 Programming,
	"Limbo":                    Programming,
	"Linker Script":            Data,
	"Linux Kernel Module":      Data,
	"Liquid":                   Markup,
	"Literate Agda":            Programming,
	"Literate CoffeeScript":    Programming,
	"Literate Haskell":         Programming,
	"LiveScript":               Programming,
	"Logos":                    Programming,
	"Logtalk":                  Programming,
	"LookML":                   Programming,
	"LoomScript":               Programming,
	"Lua":                      Programming,
	"M":                        Programming,
	"M4":                       Programming,
	"M4Sugar":                  Programming,
	"MAXScript":                Programming,
	"MQL4":                     Programming,
	"MQL5":                     Programming,
	"MTML":                     Markup,
	"MUF":                      Programming,
	"Makefile":                 Programming,
	"Mako":                     Programming,
	"Markdown":                 Prose,
	"Marko":                    Markup,
	"Mask":                     Markup,
	"Mathematica":              Programming,
	"Matlab":                   Programming,
	"Maven POM":                Data,
	"Max":                      Programming,
	"MediaWiki":                Prose,
	"Mercury":                  Programming,
	"Meson":                    Programming,
	"Metal":                    Programming,
	"MiniD":                    Programming,
	"Mirah":                    Programming,
	"Modelica":                 Programming,
	"Modula-2":                 Programming,
	"Module Management System": Programming,
	"Monkey":                   Programming,
	"Moocode":                  Programming,
	"MoonScript":               Programming,
	"Myghty":                   Programming,
	"NCL":                      Programming,
	"NL":                       Data,
	"NSIS":                     Programming,
	"Nemerle":                  Programming,
	"NetLinx":                  Programming,
	"NetLinx+ERB":              Programming,
	"NetLogo":                  Programming,
	"NewLisp":                  Programming,
	"Nginx":                    Markup,
	"Nim":                      Programming,
	"Ninja":                    Data,
	"Nit":                      Programming,
	"Nix":                      Programming,
	"Nu":                       Programming,
	"NumPy":                    Programming,
	"OCaml":                    Programming,
	"ObjDump":                  Data,
	"Objective-C":              Programming,
	"Objective-C++":            Programming,
	"Objective-J":              Programming,
	"Omgrofl":                  Programming,
	"Opa":                      Programming,
	"Opal":                     Programming,
	"OpenCL":                   Programming,
	"OpenEdge ABL":             Programming,
	"OpenRC runscript":         Programming,
	"OpenSCAD":                 Programming,
	"OpenType Feature File":    Data,
	"Org":                            Prose,
	"Ox":                             Programming,
	"Oxygene":                        Programming,
	"Oz":                             Programming,
	"P4":                             Programming,
	"PAWN":                           Programming,
	"PHP":                            Programming,
	"PLSQL":                          Programming,
	"PLpgSQL":                        Programming,
	"POV-Ray SDL":                    Programming,
	"Pan":                            Programming,
	"Papyrus":                        Programming,
	"Parrot":                         Programming,
	"Parrot Assembly":                Programming,
	"Parrot Internal Representation": Programming,
	"Pascal":                       Programming,
	"Pep8":                         Programming,
	"Perl":                         Programming,
	"Perl6":                        Programming,
	"Pic":                          Markup,
	"Pickle":                       Data,
	"PicoLisp":                     Programming,
	"PigLatin":                     Programming,
	"Pike":                         Programming,
	"Pod":                          Prose,
	"PogoScript":                   Programming,
	"Pony":                         Programming,
	"PostScript":                   Markup,
	"PowerBuilder":                 Programming,
	"PowerShell":                   Programming,
	"Processing":                   Programming,
	"Prolog":                       Programming,
	"Propeller Spin":               Programming,
	"Protocol Buffer":              Markup,
	"Public Key":                   Data,
	"Pug":                          Markup,
	"Puppet":                       Programming,
	"Pure Data":                    Programming,
	"PureBasic":                    Programming,
	"PureScript":                   Programming,
	"Python":                       Programming,
	"Python console":               Programming,
	"Python traceback":             Data,
	"QML":                          Programming,
	"QMake":                        Programming,
	"R":                            Programming,
	"RAML":                         Markup,
	"RDoc":                         Prose,
	"REALbasic":                    Programming,
	"REXX":                         Programming,
	"RHTML":                        Markup,
	"RMarkdown":                    Prose,
	"RPM Spec":                     Data,
	"RUNOFF":                       Markup,
	"Racket":                       Programming,
	"Ragel":                        Programming,
	"Rascal":                       Programming,
	"Raw token data":               Data,
	"Reason":                       Programming,
	"Rebol":                        Programming,
	"Red":                          Programming,
	"Redcode":                      Programming,
	"Regular Expression":           Data,
	"Ren'Py":                       Programming,
	"RenderScript":                 Programming,
	"RobotFramework":               Programming,
	"Roff":                         Markup,
	"Rouge":                        Programming,
	"Ruby":                         Programming,
	"Rust":                         Programming,
	"SAS":                          Programming,
	"SCSS":                         Markup,
	"SMT":                          Programming,
	"SPARQL":                       Data,
	"SQF":                          Programming,
	"SQL":                          Data,
	"SQLPL":                        Programming,
	"SRecode Template":             Markup,
	"STON":                         Data,
	"SVG":                          Data,
	"Sage":                         Programming,
	"SaltStack":                    Programming,
	"Sass":                         Markup,
	"Scala":                        Programming,
	"Scaml":                        Markup,
	"Scheme":                       Programming,
	"Scilab":                       Programming,
	"Self":                         Programming,
	"ShaderLab":                    Programming,
	"Shell":                        Programming,
	"ShellSession":                 Programming,
	"Shen":                         Programming,
	"Slash":                        Programming,
	"Slim":                         Markup,
	"Smali":                        Programming,
	"Smalltalk":                    Programming,
	"Smarty":                       Programming,
	"SourcePawn":                   Programming,
	"Spline Font Database":         Data,
	"Squirrel":                     Programming,
	"Stan":                         Programming,
	"Standard ML":                  Programming,
	"Stata":                        Programming,
	"Stylus":                       Markup,
	"SubRip Text":                  Data,
	"Sublime Text Config":          Data,
	"SuperCollider":                Programming,
	"Swift":                        Programming,
	"SystemVerilog":                Programming,
	"TI Program":                   Programming,
	"TLA":                          Programming,
	"TOML":                         Data,
	"TXL":                          Programming,
	"Tcl":                          Programming,
	"Tcsh":                         Programming,
	"TeX":                          Markup,
	"Tea":                          Markup,
	"Terra":                        Programming,
	"Text":                         Prose,
	"Textile":                      Prose,
	"Thrift":                       Programming,
	"Turing":                       Programming,
	"Turtle":                       Data,
	"Twig":                         Markup,
	"Type Language":                Data,
	"TypeScript":                   Programming,
	"Unified Parallel C":           Programming,
	"Unity3D Asset":                Data,
	"Unix Assembly":                Programming,
	"Uno":                          Programming,
	"UnrealScript":                 Programming,
	"UrWeb":                        Programming,
	"VCL":                          Programming,
	"VHDL":                         Programming,
	"Vala":                         Programming,
	"Verilog":                      Programming,
	"Vim script":                   Programming,
	"Visual Basic":                 Programming,
	"Volt":                         Programming,
	"Vue":                          Markup,
	"Wavefront Material":           Data,
	"Wavefront Object":             Data,
	"Web Ontology Language":        Markup,
	"WebAssembly":                  Programming,
	"WebIDL":                       Programming,
	"World of Warcraft Addon Data": Data,
	"X10":              Programming,
	"XC":               Programming,
	"XCompose":         Data,
	"XML":              Data,
	"XPages":           Programming,
	"XProc":            Programming,
	"XQuery":           Programming,
	"XS":               Programming,
	"XSLT":             Programming,
	"Xojo":             Programming,
	"Xtend":            Programming,
	"YAML":             Data,
	"YANG":             Data,
	"Yacc":             Programming,
	"Zephir":           Programming,
	"Zimpl":            Programming,
	"desktop":          Data,
	"eC":               Programming,
	"edn":              Data,
	"fish":             Programming,
	"mupad":            Programming,
	"nesC":             Programming,
	"ooc":              Programming,
	"reStructuredText": Prose,
	"wisp":             Programming,
	"xBase":            Programming,
}
