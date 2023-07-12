//Programm checks the usefulness of the backup PAW content and copy to AWS s3
//Created by Mikhail Kirillov 12.06.2023

package main

import (	
	"fmt"
	"log"
	"os"
	"os/exec"	
	"time"
)

//The Function checks the usefulness of the backup
func backup(time_ string) {	
	a, _ := os.ReadDir("/data/tmp/backup")
	if len(a) < 4 {
		fmt.Println(time_, " Without_redis\n")
		time.Sleep(600 * time.Second)
		main()
	} else {
		fmt.Println(time_, " Full_backup\n")
		archive(time_)
	}			
}

//The Function archive full backup
func archive(time_ string){	
	out, err := exec.Command("bash", "-c", "/usr/bin/zip -r /data/tmp/go_prog/go_archive-paw-uat-$(date +'%Y-%m-%d').zip /data/tmp/backup").Output()
	if err != nil {
        fmt.Println("%s", err)
    	}
    	fmt.Println(time_, "Archive completed successfully\n")
    	output := string(out[:])
    	fmt.Println(output)	
		delete(time_)	
}

//The Function delete old backup before new
func delete(time_ string) {
	out, err := exec.Command("bash", "-c", "rm -rf /data/tmp/backup").Output()
	if err != nil {
        fmt.Println("%s", err)
    	}
    	fmt.Println(time_, "The backup folder was successfully deleted\n")
		output := string(out[:])
    	fmt.Println(output)		
		to_aws(time_)
}

//The Function copy archive to AWS s3
func to_aws(time_ string){
	out, err := exec.Command("bash", "-c", "/usr/local/bin/aws s3 cp /data/tmp/go_prog/go_archive-paw-uat-$(date +'%Y-%m-%d').zip s3://tkio-uat-defra-ibmpa-archive-00/Backups/PAW/").Output()
	if err != nil {
		fmt.Println("%s", err)
		}
		fmt.Println(time_, "Archive was successfully copied to AWS")
		output := string(out[:])
		fmt.Println(output)
		os.Exit(0)
}

//Main Function, execute PAW backup script(Start program)
func main() {
	time_ := time.Now()
	out := exec.Command("bash", "-c", "/data/PAWD_2054/tools/backup.sh /data/tmp/backup")
		err := out.Start()
	logFile, err := os.OpenFile("Backup_PAW.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
		//output := string(out[:])
        fmt.Println(time_, out)
	log.SetOutput(logFile)
		err = out.Wait()
	backup(time_.String())
}
