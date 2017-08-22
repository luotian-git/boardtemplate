package main

import "fmt"
import "flag"

import "html/template"

//import "text/template"
import "os"
import "model"

func dockerfile() {
	myimage := model.Dockerfile{Base: "luotian1/learn-ping",
		CopyFrom: "*",
		CopyTo:   "/tmp/",
		Command:  "ping 10.110.18.70"}

	//fmt.Println("dockerfile = ", myimage)
	var tmpl *template.Template
	//tmpl := template.New("dockerfile")
	//tmpl.Parse("FROM {{.Base}}\nCOPY {{.CopyFrom}} {{.CopyTo}}\nENTRYPOINT {{.Entrypoint}} \n")
	tmpl, err := template.ParseFiles("dockerfile-template.txt")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(os.Stdout, myimage)
	fmt.Println("")
}

func deployment() {
	//define a data structure for testing
	myenv := model.ContainerEnv{"name", "value"}
	myvolume := model.ContainerVolume{"name", "path"}
	mycontainer := model.Container{
		ContainerName:  "bigdata",
		BaseImage:      "bigdata",
		ContainerPorts: []int{8080},
		VolumeList:     []model.ContainerVolume{myvolume},
		EnvList:        []model.ContainerEnv{myenv}}
	mydeployment := model.Deployment{
		DeploymentName: "bigdata",
		Replicas:       1,
		NFSvolumeList:  []model.NFSvolume{{"name1", "server", "path"}},
		ContainerList:  []model.Container{mycontainer}}

	//fmt.Println("deployment = ", mydeployment)
	var tmpl *template.Template
	tmpl, err := template.ParseFiles("deployment-template")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(os.Stdout, mydeployment)
	fmt.Println("")
}

func deployment2() {
	//define a bigdata deployment
	myenv := model.ContainerEnv{"MYSQL_ROOT_PASSWORD", "123456Aa?"}
	myenvs2 := []model.ContainerEnv{{"DATASOURCE_DRIVERCLASSNAME", "org.gjt.mm.mysql.Driver"},
		{"DATASOURCE_URL", "jdbc:mysql://localhost:3306/idap?useUnicode=true&characterEncoding=utf-8&allowMultiQueries=true&useOldAliasMetadataBehavior=true"},
		{"DATASOURCE_USERNAME", "root"},
		{"DATASOURCE_PASSWORD", " 123456Aa?"},
		{"BIGDATA_DOMAIN", "10.110.18.58"}}
	//myvolume := model.ContainerVolume{"name", "path"}
	mycontainer1 := model.Container{
		ContainerName:  "bigdatamysql",
		BaseImage:      "10.110.13.58:5000/bigdata-mysql:latest",
		ContainerPorts: []int{3306},
		//VolumeList:     []model.ContainerVolume{myvolume},
		EnvList: []model.ContainerEnv{myenv}}
	mycontainer2 := model.Container{
		ContainerName:  "bigdatatomcat",
		BaseImage:      "10.110.13.58:5000/bigdata-tomcat:latest",
		ContainerPorts: []int{8080},
		//VolumeList:     []model.ContainerVolume{myvolume},
		EnvList: myenvs2}
	mydeployment := model.Deployment{
		DeploymentName: "bigdata",
		Replicas:       1,
		//NFSvolumeList:  []model.NFSvolume{{"name1", "server", "path"}},
		ContainerList: []model.Container{mycontainer1, mycontainer2}}

	//fmt.Println("deployment = ", mydeployment)
	var tmpl *template.Template
	tmpl, err := template.ParseFiles("deployment-template")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(os.Stdout, mydeployment)
	fmt.Println("")
}

func service() {
	//define a data structure for testing
	myport := model.ServicePort{8080, 8080, 8081}
	myservice := model.Service{
		ServiceName:     "bigdata",
		ExternalService: true,
		PortList:        []model.ServicePort{myport},
		SelectorList:    []string{"bigdata"}}

	//fmt.Println("service = ", myservice)
	var tmpl *template.Template
	tmpl, err := template.ParseFiles("service-template")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(os.Stdout, myservice)
	fmt.Println("")
}

func service2() {
	//define a bigdata service structure for testing
	myport := model.ServicePort{8080, 8080, 30009}
	myservice := model.Service{
		ServiceName:     "bigdata",
		ExternalService: true,
		PortList:        []model.ServicePort{myport},
		SelectorList:    []string{"bigdata"}}

	//fmt.Println("service = ", myservice)
	var tmpl *template.Template
	tmpl, err := template.ParseFiles("service-template")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(os.Stdout, myservice)
	fmt.Println("")
}

func main() {
	templateName := flag.String("template", "dockerfile", "template name")
	flag.Parse()
	switch *templateName {
	case "dockerfile":
		dockerfile()
	case "deployment":
		deployment2()
	case "service":
		service2()
	}

}
