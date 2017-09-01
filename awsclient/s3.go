package awsclient

import "github.com/aws/aws-sdk-go/service/s3"

type S3 interface {
	PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

// {
//     "schemaVersion":"2.2",
//     "description":"Run a shell script or specify the path to a script to run.",
//     "parameters":{
//         "commands":{
//             "type":"StringList",
//             "description":"(Required) Specify the commands to run or the paths to existing scripts on the instance.",
//             "minItems":1,
//             "displayType":"textarea"
//         },
//         "workingDirectory":{
//             "type":"String",
//             "default":"",
//             "description":"(Optional) The path to the working directory on your instance.",
//             "maxChars":4096
//         },
//         "executionTimeout":{
//             "type":"String",
//             "default":"3600",
//             "description":"(Optional) The time in seconds for a command to be completed before it is considered to have failed. Default is 3600 (1 hour). Maximum is 28800 (8 hours).",
//             "allowedPattern":"([1-9][0-9]{0,3})|(1[0-9]{1,4})|(2[0-7][0-9]{1,3})|(28[0-7][0-9]{1,2})|(28800)"
//         }
//     },
//        "mainSteps":[
//       {
//          "action":"aws:runShellScript",
//          "name":"a",
//          "inputs":{
//                     "runCommand":"{{ commands }}",
//                     "workingDirectory":"{{ workingDirectory }}",
//                     "timeoutSeconds":"{{ executionTimeout }}"

//          }
//       }
//    ]
// }
