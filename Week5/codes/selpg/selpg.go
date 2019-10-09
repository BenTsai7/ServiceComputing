package main

/*================================= includes ======================*/

import (
	"os"
	"os/exec"
	"fmt"
	"io"
	"bufio"
	 flag "github.com/spf13/pflag"
)

/*================================= types =========================*/

type selpg_args struct{
	start_page int
	end_page int
	in_filename string
	page_len int /* default value, can be overriden by "-l number" on command line */
	page_type int/* 'l' for lines-delimited, 'f' for form-feed-delimited */
					/* default is 'l' */
	print_dest string
}

type sp_args selpg_args;

/*================================= globals =======================*/

var progname string/* program name, for error messages */
var fin *os.File
var fout *os.File

/*================================= prototypes ====================*/

// GOLANG FUNCTION PROTOTYPES NOT SUPPORT
//func usage(void)
//func process_args(psa *sp_args)
//func process_input(sp_args sa)

/*================================= main()=== =====================*/
func main(){
	var sa sp_args
	progname = os.Args[0];
	sa.start_page = -1
	sa.end_page = -1
	sa.page_len = 72
	sa.page_type = 'l'

	process_args(&sa);
	process_input(sa);
	fmt.Println(progname);
}

/*================================= process_args() ================*/
func process_args(psa *sp_args){
	//use pflag to parse 
	var start_page = flag.IntP("s", "s", -1, "start page ")
	var end_page = flag.IntP("e", "e", -1, "end page")
	var l = flag.IntP("l", "l", 72, "page length")
	var f = flag.BoolP("f","f", false, "page seperator")
	var d = flag.StringP("d","d","","destination")
	flag.Parse()

	if(*start_page==-1 || *end_page==-1){
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", progname)
		usage()
		os.Exit(1)
	}
	const INT_MAX = int(^uint(0) >> 1)
	if (*start_page < 1 || *start_page >(INT_MAX - 1)){
		fmt.Fprintf(os.Stderr, "%s: invalid start page %d\n", progname, *start_page)
		usage()
		os.Exit(2)
	}
	(*psa).start_page = *start_page

	if ((*end_page < 1) || (*end_page > (INT_MAX - 1)) || (*end_page < *start_page)){
		fmt.Fprintf(os.Stderr, "%s: invalid end page %d\n", progname, *end_page)
		usage();
		os.Exit(3)
	}
	(*psa).end_page = *end_page
	if (*l < 1 || *l >(INT_MAX - 1)){
		fmt.Fprintf(os.Stderr, "%s: invalid page length %d\n", progname, *l)
		usage()
		os.Exit(4)
	}
	(*psa).page_type = 'l'
	(*psa).page_len = *l
	if (*f){
		(*psa).page_type = 'f'
	}
	if (*d!=""){
		(*psa).print_dest = *d;
	}
	// using os to check filename args
	if (os.Args[len(os.Args)-1][0]!='-'){
		(*psa).in_filename = os.Args[len(os.Args)-1]
		_, err := os.Stat((*psa).in_filename)
		if(err!=nil){
			fmt.Fprintf(os.Stderr, "%s: input file \"%s\" does not exist\n",
				progname, (*psa).in_filename)
			os.Exit(5)
		}		
	}
}

/*================================= process_input() ===============*/

func process_input(sa sp_args){
	var (
		stdin io.WriteCloser
		cmd *exec.Cmd
		line_ctr int /* line counter */
		page_ctr int /* page counter */
		err error
		line []byte
		stringline string
	)
	if (sa.in_filename == ""){
		fin = os.Stdin
	}else{
		fin, err = os.Open(sa.in_filename)
		if(err != nil){
			fmt.Fprintf(os.Stderr, "%s: could not open input file \"%s\"\n",
				progname, sa.in_filename);
			os.Exit(6)
		}
		defer fin.Close()
	}
	if (sa.print_dest == ""){
		fout = os.Stdout
	}else{
		//cmd = exec.Command("cat", "-n")
		cmd = exec.Command("./printer")
		stdin, err = cmd.StdinPipe()
		if(err != nil){
			fmt.Fprintf(os.Stderr, "%s: could not open pipe to \"%s\"\n",
				progname, sa.print_dest)
			os.Exit(7)
		}
		defer fout.Close()
	}
	reader := bufio.NewReader(fin)
	writer := bufio.NewWriter(fout)
	
	if (sa.page_type == 'l'){
		line_ctr = 0
		page_ctr = 1
		for {	
			line, _, err = reader.ReadLine()
			if(err!=nil && err!=io.EOF){
				fmt.Fprintf(os.Stderr, "%s: system error [%s] occurred on input stream fin\n",
				progname, sa.in_filename)
				os.Exit(14)
			}
			if (err == io.EOF){ /* error or EOF */
				break
			}
			line_ctr++;
			if (line_ctr > sa.page_len){
				page_ctr++
				line_ctr = 1
			}
			if ((page_ctr >= sa.start_page) && (page_ctr <= sa.end_page)){
				if(sa.print_dest != ""){
					stdin.Write(line)
					stdin.Write([]byte("\n"))
				}else{
					writer.Write(line)
					writer.WriteString("\n")
				}
			}
		}
	}else{
		page_ctr = 1
		for{
			stringline, err = reader.ReadString('\f')
			if (err == io.EOF){ /* error or EOF */
				break
			}
			page_ctr++;
			if ((page_ctr >= sa.start_page) && (page_ctr <= sa.end_page)){
				if(sa.print_dest != ""){
					stdin.Write([]byte(stringline))
				}else{
					writer.Write([]byte(stringline))
				}
				
			}
		}
	}
	
	/* end main loop */

	if (page_ctr < sa.start_page){
		fmt.Fprintf(os.Stderr,
			"%s: start_page (%d) greater than total pages (%d), no output written\n", progname, sa.start_page, page_ctr)
	}else if (page_ctr < sa.end_page){
		fmt.Fprintf(os.Stderr, 
			"%s: end_page (%d) greater than total pages (%d), less output than expected\n", progname, sa.end_page, page_ctr)
	}
	writer.Flush()
	fmt.Fprintf(os.Stderr, "%s: done\n", progname)
	if sa.print_dest != "" {
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

/*================================= usage() =======================*/

func usage(){
	fmt.Fprintf(os.Stderr,
		"\nUSAGE: %s -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]\n", progname);
}

/*================================= EOF ===========================*/
