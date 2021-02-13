package crazyexec

import "fmt"
import "time"
import "os/exec"
import "os"
import "./text"
import "./cinstall"
//######### VARIABLES
var pause,errfile,subdata string
var (
	rojo string = "\033[31m"
	verde string = "\033[32m"
	amarillo string = "\033[33m"
	azul string = "\033[36m"
	blanco string = "\033[37m"
	fnegro string = "\033[40m"
	negro string = "\033[30m"
	famarillo string = "\033[43m"
)
// ######## EJECUSIONES PRINCIPALES
func Bienvenida(){
	clear()
	fmt.Printf("    %v _______________________________\n",blanco)
	fmt.Printf("    %v|%v%v_CrazyPop_____________[_][-][X]%v%v|\n",blanco,negro,famarillo,fnegro,blanco)
	fmt.Printf("    %v|%v%v Iniciando CrazyPop (^o^)  1.3 %v%v|\n",blanco,negro,famarillo,fnegro,blanco)
	fmt.Printf("    %v|%v%v        BIENVENIDO !!!         %v%v|\n",blanco,negro,famarillo,fnegro,blanco)
	fmt.Printf("    %v|%v%v_______________________________%v%v|\n\n",blanco,negro,famarillo,fnegro,blanco)
	time.Sleep(5*time.Second)
}
// ####### Instalaciones de software
func Nmap(){
	for r := 1; r >= 1; r++{
		clear();text.Nmap();subdata = " "
		fmt.Printf("CrazyPop/Nmap: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Nmap()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Yow(){
	for r := 1; r >= 1; r++{
		clear();text.Yow();subdata = " "
		fmt.Printf("CrazyPop/Yow: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Yow()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Youtubedl(){
	for r := 1; r >= 1; r++{
		clear();text.Youtubedl();subdata = " "
		fmt.Printf("CrazyPop/Youtubedl: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Youtubedl()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Moonbuggy(){
	for r := 1; r >= 1; r++{
		clear();text.Moonbuggy();subdata = " "
		fmt.Printf("CrazyPop/Moon-Buggy: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Moonbuggy()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Clamav(){
	for r := 1; r >= 1; r++{
		clear();text.Clamav();subdata = " "
		fmt.Printf("CrazyPop/ClamAV: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Clamav()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Hydra(){
	for r := 1; r >= 1; r++{
		clear();text.Hydra();subdata = " "
		fmt.Printf("CrazyPop/Hydra: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Hydra()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Ubuntu(){
	for r := 1; r >= 1; r++{
		clear();text.Ubuntu();subdata = " "
		fmt.Printf("CrazyPop/Ubuntu: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Ubuntu()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Kali(){
	for r := 1; r >= 1; r++{
		clear();text.Kali();subdata = " "
		fmt.Printf("CrazyPop/Kali: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Kali()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Pacman(){
	for r := 1; r >= 1; r++{
		clear();text.Pacman();subdata = " "
		fmt.Printf("CrazyPop/PacMan: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Pacman()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Freezer(){
	for r := 1; r >= 1; r++{
		clear();text.Freezer();subdata = " "
		fmt.Printf("CrazyPop/Freezer: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Freezer()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Darksploit(){
	for r := 1; r >= 1; r++{
		clear();text.Darksploit();subdata = " "
		fmt.Printf("CrazyPop/DarkSploit: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Darksploit()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Tor(){
	for r := 1; r >= 1; r++{
		clear();text.Tor();subdata = " "
		fmt.Printf("CrazyPop/Tor: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Tor()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Weeman(){
	for r := 1; r >= 1; r++{
		clear();text.Weeman();subdata = " "
		fmt.Printf("CrazyPop/Weeman: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Weeman()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Crunch(){
	for r := 1; r >= 1; r++{
		clear();text.Crunch();subdata = " "
		fmt.Printf("CrazyPop/Crunch: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Crunch()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Debian(){
	for r := 1; r >= 1; r++{
		clear();text.Debian();subdata = " "
		fmt.Printf("CrazyPop/Debian: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Debian()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Exiftool(){
	for r := 1; r >= 1; r++{
		clear();text.Exiftool();subdata = " "
		fmt.Printf("CrazyPop/Exiftool: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Exiftool()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Gnupg(){
	for r := 1; r >= 1; r++{
		clear();text.Gnupg();subdata = " "
		fmt.Printf("CrazyPop/Gnupg: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Gnupg()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Cupp(){
	for r := 1; r >= 1; r++{
		clear();text.Cupp();subdata = " "
		fmt.Printf("CrazyPop/Cupp: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Cupp()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Infoga(){
	for r := 1; r >= 1; r++{
		clear();text.Infoga();subdata = " "
		fmt.Printf("CrazyPop/Infoga: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Cupp()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Phoneinfoga(){
	for r := 1; r >= 1; r++{
		clear();text.Phoneinfoga();subdata = " "
		fmt.Printf("CrazyPop/Phoneinfoga: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Phoneinfoga()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Tmux(){
	for r := 1; r >= 1; r++{
		clear();text.Tmux();subdata = " "
		fmt.Printf("CrazyPop/Tmux: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Tmux()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Torshammer(){
	for r := 1; r >= 1; r++{
		clear();text.Torshammer();subdata = " "
		fmt.Printf("CrazyPop/Torshammer: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Torshammer()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Slowloris(){
	for r := 1; r >= 1; r++{
		clear();text.Slowloris();subdata = " "
		fmt.Printf("CrazyPop/Slowloris: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Slowloris()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Spiderbot(){
	for r := 1; r >= 1; r++{
		clear();text.Spiderbot();subdata = " "
		fmt.Printf("CrazyPop/Spiderbot: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Spiderbot()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Sqlmap(){
	for r := 1; r >= 1; r++{
		clear();text.Sqlmap();subdata = " "
		fmt.Printf("CrazyPop/SQLmap: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Sqlmap()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Cmatrix(){
	for r := 1; r >= 1; r++{
		clear();text.Cmatrix();subdata = " "
		fmt.Printf("CrazyPop/Cmatrix: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Cmatrix()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func W3m(){
	for r := 1; r >= 1; r++{
		clear();text.W3m();subdata = " "
		fmt.Printf("CrazyPop/W3m: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.W3m()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Htop(){
	for r := 1; r >= 1; r++{
		clear();text.Htop();subdata = " "
		fmt.Printf("CrazyPop/Htop: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Htop()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Hashcat(){
	for r := 1; r >= 1; r++{
		clear();text.Hashcat();subdata = " "
		fmt.Printf("CrazyPop/Hashcat: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Hashcat()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Lazymux(){
	for r := 1; r >= 1; r++{
		clear();text.Lazymux();subdata = " "
		fmt.Printf("CrazyPop/Lazymux: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Lazymux()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Wpspin(){
	for r := 1; r >= 1; r++{
		clear();text.Wpspin();subdata = " "
		fmt.Printf("CrazyPop/Wpspin: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Wpspin()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Commix(){
	for r := 1; r >= 1; r++{
		clear();text.Commix();subdata = " "
		fmt.Printf("CrazyPop/Commix: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Commix()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Sms(){
	for r := 1; r >= 1; r++{
		clear();text.Sms();subdata = " "
		fmt.Printf("CrazyPop/Spazsms: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Sms()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Neofetch(){
	for r := 1; r >= 1; r++{
		clear();text.Neofetch();subdata = " "
		fmt.Printf("CrazyPop/Neofetch: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Neofetch()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Flappy(){
	for r := 1; r >= 1; r++{
		clear();text.Flappy();subdata = " "
		fmt.Printf("CrazyPop/Flappy_Bird: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Flappy()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Vbug(){
	for r := 1; r >= 1; r++{
		clear();text.Vbug();subdata = " "
		fmt.Printf("CrazyPop/Vbug: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Vbug()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Routersploit(){
	for r := 1; r >= 1; r++{
		clear();text.Routersploit();subdata = " "
		fmt.Printf("CrazyPop/RouterSploit: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Routersploit()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Blackhydra(){
	for r := 1; r >= 1; r++{
		clear();text.Blackhydra();subdata = " "
		fmt.Printf("CrazyPop/Black_Hydra: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Blackhydra()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Websploit(){
	for r := 1; r >= 1; r++{
		clear();text.Websploit();subdata = " "
		fmt.Printf("CrazyPop/WebSploit: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Websploit()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Astranmap(){
	for r := 1; r >= 1; r++{
		clear();text.Astranmap();subdata = " "
		fmt.Printf("CrazyPop/AstraNmap: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Astranmap()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
func Arat(){
	for r := 1; r >= 1; r++{
		clear();text.Arat();subdata = " "
		fmt.Printf("CrazyPop/A-Rat: ")
		fmt.Scanf("%v",&subdata)
		if subdata == "2"{
			salir()
		}else if subdata == "1"{
			cinstall.Arat()
		}else if subdata == "3"{
			break
		}else{
			fmt.Printf("Error en (¬!^-^)¬ (%v)",subdata)
			time.Sleep(1*time.Second)
		}
	}
}
//############## EJECUSIONES SECUNDARIOS

func clear(){ //Limpiador de pantalla
	cl := exec.Command("clear")
	cl.Stdout = os.Stdout
	cl.Run()
}
func pausa(){  //Pausa del programa
	time.Sleep(1*time.Second)
	fmt.Printf("\n\n%v(^o^) %vENTER PARA CONTINUAR%v",verde,amarillo,blanco)
	fmt.Scanf("%v",&pause)
}
func salir(){ // Salida del programa
	fmt.Printf("%vSaliendo %v(T-T)%v\n",rojo,verde,blanco)
	time.Sleep(1*time.Second)
	clear()
	fmt.Printf("%vBye Bye%v\n",azul,blanco)
	os.Exit(0)
}
