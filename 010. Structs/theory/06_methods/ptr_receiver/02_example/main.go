package main

import "fmt"

type Server struct {
	// Операционная система
	// Правило: Ubuntu
	OperationSystem string

	// Кол-во ядер ЦПУ
	// Правило: не меньше 4
	CpuNum int

	// Кол-во РАМ
	// Правило: не меньше 4 000 МБ
	RamValue int

	// Размер диска
	// Правило:  не меньше 20 000 МБ
	StorageValue int

	// Какой сервис должен быть установлен
	// Правило: nginx
	ServiceInstall string

	// Доступность сервера в/из интернет
	// Правило: должен
	InternetAccess bool
}

func (s *Server) ChangeOS(newOS string) {
	if newOS != "Ubuntu" {
		fmt.Println("We need to use only Ubuntu")
	}
}

func (s *Server) ChangeCpuNum(newCpuNumValue int) {
	if newCpuNumValue <= 4 {
		fmt.Println("We need to use at least 4 CPU Core")
		fmt.Println("The selected value is:", newCpuNumValue)
	} else {
		s.CpuNum = newCpuNumValue
		fmt.Println("I can increase CPU Cores.")
		fmt.Println("New value is:", newCpuNumValue)
	}
}

func main() {
	server := Server{
		OperationSystem: "Ubuntu",
		CpuNum:          4,
		RamValue:        3500,
		StorageValue:    25000,
		ServiceInstall:  "nginx",
		InternetAccess:  true,
	}
	fmt.Println("I need server with next chars:", server)
	fmt.Println()
	fmt.Println("Current CPU Num Cores:", server.CpuNum)
	server.ChangeCpuNum(5)
	server.ChangeOS("Debian")
	
}
