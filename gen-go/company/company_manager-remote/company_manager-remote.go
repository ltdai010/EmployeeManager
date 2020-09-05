// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"github.com/apache/thrift/lib/go/thrift"
	"company"
)

var _ = company.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  Employee getEmployee(string id, string companyID)")
  fmt.Fprintln(os.Stderr, "  void postEmployee(string id, string name, string address, Date date, string company)")
  fmt.Fprintln(os.Stderr, "  void putEmployee(string id, string name, string address, Date date, string company)")
  fmt.Fprintln(os.Stderr, "  void removeEmployee(string id, string companyID)")
  fmt.Fprintln(os.Stderr, "  Company getCompany(string id)")
  fmt.Fprintln(os.Stderr, "   getAllEmployee()")
  fmt.Fprintln(os.Stderr, "   getListEmployee(string companyID, int start, int count)")
  fmt.Fprintln(os.Stderr, "   getListEmployeeInDate(string companyID, Date pros, Date cons)")
  fmt.Fprintln(os.Stderr, "   getAllCompany()")
  fmt.Fprintln(os.Stderr, "  void postCompany(string id, string name, string address)")
  fmt.Fprintln(os.Stderr, "  void putCompany(string id, string name, string address)")
  fmt.Fprintln(os.Stderr, "   getEmployeeList(string id)")
  fmt.Fprintln(os.Stderr, "  void removeCompany(string id)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := company.NewCompanyManagerClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "getEmployee":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetEmployee requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.GetEmployee(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "postEmployee":
    if flag.NArg() - 1 != 5 {
      fmt.Fprintln(os.Stderr, "PostEmployee requires 5 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    arg38 := flag.Arg(4)
    mbTrans39 := thrift.NewTMemoryBufferLen(len(arg38))
    defer mbTrans39.Close()
    _, err40 := mbTrans39.WriteString(arg38)
    if err40 != nil {
      Usage()
      return
    }
    factory41 := thrift.NewTJSONProtocolFactory()
    jsProt42 := factory41.GetProtocol(mbTrans39)
    argvalue3 := company.NewDate()
    err43 := argvalue3.Read(jsProt42)
    if err43 != nil {
      Usage()
      return
    }
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    fmt.Print(client.PostEmployee(context.Background(), value0, value1, value2, value3, value4))
    fmt.Print("\n")
    break
  case "putEmployee":
    if flag.NArg() - 1 != 5 {
      fmt.Fprintln(os.Stderr, "PutEmployee requires 5 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    arg48 := flag.Arg(4)
    mbTrans49 := thrift.NewTMemoryBufferLen(len(arg48))
    defer mbTrans49.Close()
    _, err50 := mbTrans49.WriteString(arg48)
    if err50 != nil {
      Usage()
      return
    }
    factory51 := thrift.NewTJSONProtocolFactory()
    jsProt52 := factory51.GetProtocol(mbTrans49)
    argvalue3 := company.NewDate()
    err53 := argvalue3.Read(jsProt52)
    if err53 != nil {
      Usage()
      return
    }
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    fmt.Print(client.PutEmployee(context.Background(), value0, value1, value2, value3, value4))
    fmt.Print("\n")
    break
  case "removeEmployee":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "RemoveEmployee requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.RemoveEmployee(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "getCompany":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetCompany requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetCompany(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getAllEmployee":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetAllEmployee requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetAllEmployee(context.Background()))
    fmt.Print("\n")
    break
  case "getListEmployee":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "GetListEmployee requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err59 := (strconv.Atoi(flag.Arg(2)))
    if err59 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := company.Int(argvalue1)
    tmp2, err60 := (strconv.Atoi(flag.Arg(3)))
    if err60 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := company.Int(argvalue2)
    fmt.Print(client.GetListEmployee(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "getListEmployeeInDate":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "GetListEmployeeInDate requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg62 := flag.Arg(2)
    mbTrans63 := thrift.NewTMemoryBufferLen(len(arg62))
    defer mbTrans63.Close()
    _, err64 := mbTrans63.WriteString(arg62)
    if err64 != nil {
      Usage()
      return
    }
    factory65 := thrift.NewTJSONProtocolFactory()
    jsProt66 := factory65.GetProtocol(mbTrans63)
    argvalue1 := company.NewDate()
    err67 := argvalue1.Read(jsProt66)
    if err67 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    arg68 := flag.Arg(3)
    mbTrans69 := thrift.NewTMemoryBufferLen(len(arg68))
    defer mbTrans69.Close()
    _, err70 := mbTrans69.WriteString(arg68)
    if err70 != nil {
      Usage()
      return
    }
    factory71 := thrift.NewTJSONProtocolFactory()
    jsProt72 := factory71.GetProtocol(mbTrans69)
    argvalue2 := company.NewDate()
    err73 := argvalue2.Read(jsProt72)
    if err73 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    fmt.Print(client.GetListEmployeeInDate(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "getAllCompany":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetAllCompany requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetAllCompany(context.Background()))
    fmt.Print("\n")
    break
  case "postCompany":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "PostCompany requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.PostCompany(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "putCompany":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "PutCompany requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.PutCompany(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "getEmployeeList":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetEmployeeList requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetEmployeeList(context.Background(), value0))
    fmt.Print("\n")
    break
  case "removeCompany":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RemoveCompany requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.RemoveCompany(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
