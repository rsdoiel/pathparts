//
// pathparts.go - extracts elements of a file path
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// Copyright (c) 2016, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//  list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//  this list of conditions and the following disclaimer in the documentation
//  and/or other materials provided with the distribution.
//
// * Neither the name of findfile nor the names of its
//  contributors may be used to endorse or promote products derived from
//  this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

const (
	Version = "0.0.1"
)

var (
	showHelp    bool
	showVersion bool
	showLicense bool
	Basename    bool
	Extension   bool
	Directory   bool
)

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&Basename, "b", false, "show basename")
	flag.BoolVar(&Extension, "e", false, "show extension")
	flag.BoolVar(&Directory, "d", false, "show directory")
}

func main() {
	flag.Parse()
	args := flag.Args()

	if showVersion == true {
		fmt.Printf("Version %s\n", Version)
		os.Exit(0)
	}

	if showLicense == true {
		fmt.Println(`

 Copyright (c) 2016, R. S. Doiel
 All rights reserved.

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are met:

 * Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

 * Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

 * Neither the name of findfile nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

`)
		os.Exit(0)
	}

	if showHelp == true {
		fmt.Printf(`USAGE pathparts [OPTIONS] TARGET_PATH

  Extract the path parts from options chosen.

`)
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("    -%s  (defaults to %s) %s\n", f.Name, f.DefValue, f.Usage)
		})
		fmt.Printf(" Version %s\n", Version)
		if showHelp == false && len(args) == 0 {
			os.Exit(1)
		}
		os.Exit(0)
	}

	// Handle the case where currect work directory is assumed
	if len(args) == 0 {
		args = []string{"."}
	}

	// For each directory to search walk the path
	for _, dir := range args {
		if Directory == true {
			fmt.Printf("%s ", path.Dir(dir))
		}
		if Basename == true {
			fmt.Printf("%s ", path.Base(dir))
		}
		if Extension == true {
			fmt.Printf("%s ", path.Ext(dir))
		}
		fmt.Println("")
	}
}
