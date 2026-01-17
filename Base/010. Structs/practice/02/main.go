package main

import (
	"fmt"
)

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
	// Правило: nginx по-умолчанию, но можно добавлять
	ServiceInstall []string

	// Доступность сервера в/из интернет
	// Правило: должен
	InternetAccess bool
}

func (s Server) ChangeOS(newOS string) {
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

func (s *Server) ChangeRam(newRamValue int) {
	if newRamValue < 4000 || newRamValue > 8000 {
		fmt.Println("We need to use at least 4GM RAM and no more 8GB")
		fmt.Println("The selected value is:", newRamValue)
	} else {
		s.RamValue = newRamValue
		fmt.Println("I can increase RAM.")
		fmt.Println("New value is:", newRamValue)
	}
}

func (s *Server) AddInstallService(addService string) {
	s.ServiceInstall = append(s.ServiceInstall, addService)
}

func main() {
	server := Server{
		OperationSystem: "Ubuntu",
		CpuNum:          4,
		RamValue:        4000,
		StorageValue:    25000,
		ServiceInstall:  []string{"nginx"},
		InternetAccess:  true,
	}
	fmt.Println("I need server with next chars:", server)
	fmt.Println()

	fmt.Println("Current CPU Num Cores:", server.CpuNum)
	server.ChangeCpuNum(5)
	fmt.Println()

	server.ChangeOS("Debian")
	fmt.Println()

	server.ChangeRam(5000)
	fmt.Println()

	ptr := &server
	fmt.Println("Services on server before changed:", ptr.ServiceInstall)
	server.AddInstallService("docker")
	fmt.Println("Services on server after changed:", ptr.ServiceInstall)
}
