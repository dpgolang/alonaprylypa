package main
import(
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
var file,str,newstr,newFile string
func init(){
	flag.StringVar(&file,"file","/home/alonaprylypa/go/ElementaryTasks/Task4/main/file.txt","address of file")
	flag.StringVar(&str,"old","/home/alonaprylypa/go/ElementaryTasks/Task4/main/newfile.txt","your pattern for searching")
	flag.StringVar(&newstr,"new","go","new word for replace")
	flag.StringVar(&newFile,"newfile","went","file to rewrite")
	flag.Parse()
}
func count(file,str string)(int,error){
	var count int
	data,err:=os.Open(file)
	if err!=nil{
		return 0,errors.New("file reading error")
	}
	defer data.Close()
	scanner:=bufio.NewScanner(data)
	for scanner.Scan(){
		if strings.Contains(scanner.Text(),str){
			count++
		}
	}
	if err:=scanner.Err();err!=nil{
		return 0, errors.New("counting error")
	}
	return count,nil
}
func replace(file,str,newstr,newfile string)error{
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.New("file reading error")
	}

	newData := strings.Replace(string(data), str, newstr, -1)

	err = ioutil.WriteFile(file, []byte(newData), 0)
	if err != nil {
		return errors.New("rewrite error")
	}
	return nil
}
func main(){
	number,err:=count(file,str)
	if err==nil{
		fmt.Print("Number of entries: ",number)
	}
	if replace(file,str,newstr,newFile)==nil{
		fmt.Print("no errors overwritten")
	}else{
		fmt.Print("you have some errors")
	}
}
