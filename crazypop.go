package main
import "fmt"
import "time"
import "os/exec"
import "os"
import "./crazyexec"
import "./crazytxt"
//############## Variables
var pause,data string
var rojo,verde,amarillo,blanco,azul string
//############## EJECUSION PRINCIPAL
func main(){
	rojo = "\033[31m"; verde = "\033[32m"; amarillo = "\033[33m"; blanco = "\033[37m"; azul = "\033[36m"
	crazyexec.Bienvenida()
	for a := 1; a >= 1; a++{
		clear(); data = " "
		crazytxt.Menu()
		fmt.Printf("CrazyPop: ");fmt.Scanf("%v",&data)
		if data == "salir"{
			salir()
		}else if data == "ayuda"{
			clear();crazytxt.Ayuda(); pausa()
		}else if data == "1"{
			crazyexec.Nmap()
		}else if data == "2"{
			crazyexec.Yow()
		}else if data == "3"{
			crazyexec.Youtubedl()
		}else if data == "4"{
			crazyexec.Moonbuggy()
		}else if data == "5"{
			crazyexec.Clamav()
		}else if data == "6"{
			crazyexec.Hydra()
		}else if data == "7"{
			crazyexec.Ubuntu()
		}else if data == "8"{
			crazyexec.Kali()
		}else if data == "9"{
			crazyexec.Pacman()
		}else if data == "10"{
			crazyexec.Freezer()
		}else if data == "next"{
			menu2()
		}else{
			fmt.Printf("%vError en %v(¬!^-^)¬%v (%v)",rojo,verde,blanco,data); time.Sleep(2*time.Second)
		}
	}
}
func menu2(){
	for b :=1; b>=1; b++{
		clear();data = " "
		crazytxt.Menu2()
		fmt.Printf("CrazyPop:");fmt.Scanf("%v",&data)
		if data == "salir"{
			salir()
		}else if data == "ayuda"{
			clear();crazytxt.Ayuda(); pausa()
		}else if data == "back"{
			break
		}else if data == "11"{
			crazyexec.Darksploit()
		}else if data == "12"{
			crazyexec.Tor()
		}else if data == "13"{
			crazyexec.Weeman()
		}else if data == "14"{
			crazyexec.Crunch()
		}else if data == "15"{
			crazyexec.Debian()
		}else if data == "16"{
			crazyexec.Exiftool()
		}else if data == "17"{
			crazyexec.Gnupg()
		}else if data == "18"{
			crazyexec.Cupp()
		}else if data == "19"{
			crazyexec.Infoga()
		}else if data == "20"{
			crazyexec.Phoneinfoga()
		}else if data == "next"{
			menu3()
		}else{
			fmt.Printf("%vError en %v(¬!^-^)¬%v (%v)",rojo,verde,blanco,data); time.Sleep(2*time.Second)
		}
	}
}
func menu3(){
	for b :=1; b>=1; b++{
		clear();data = " "
		crazytxt.Menu3()
		fmt.Printf("CrazyPop:");fmt.Scanf("%v",&data)
		if data == "salir"{
			salir()
		}else if data == "ayuda"{
			clear();crazytxt.Ayuda(); pausa()
		}else if data == "back"{
			break
		}else if data == "21"{
			crazyexec.Tmux()
		}else if data == "22"{
			crazyexec.Torshammer()
		}else if data == "23"{
			crazyexec.Slowloris()
		}else if data == "24"{
			crazyexec.Spiderbot()
		}else if data == "25"{
			crazyexec.Sqlmap()
		}else if data == "26"{
			crazyexec.Cmatrix()
		}else if data == "27"{
			crazyexec.W3m()
		}else if data == "28"{
			crazyexec.Htop()
		}else if data == "29"{
			crazyexec.Hashcat()
		}else if data == "30"{
			crazyexec.Lazymux()
		}else if data == "next"{
			menu4()
		}else{
			fmt.Printf("%vError en %v(¬!^-^)¬%v (%v)",rojo,verde,blanco,data); time.Sleep(2*time.Second)
		}
	}
}
func menu4(){
	for b :=1; b>=1; b++{
		clear();data = " "
		crazytxt.Menu4()
		fmt.Printf("CrazyPop:");fmt.Scanf("%v",&data)
		if data == "salir"{
			salir()
		}else if data == "ayuda"{
			clear();crazytxt.Ayuda(); pausa()
		}else if data == "back"{
			break
		}else if data == "31"{
			crazyexec.Wpspin()
		}else if data == "32"{
			crazyexec.Commix()
		}else if data == "33"{
			crazyexec.Sms()
		}else if data == "34"{
			crazyexec.Neofetch()
		}else if data == "35"{
			crazyexec.Flappy()
		}else if data == "36"{
			crazyexec.Vbug()
		}else if data == "37"{
			crazyexec.Routersploit()
		}else if data == "38"{
			crazyexec.Blackhydra()
		}else if data == "39"{
			crazyexec.Websploit()
		}else if data == "40"{
			crazyexec.Astranmap()
		}else if data == "next"{
			menu5()
		}else{
			fmt.Printf("%vError en %v(¬!^-^)¬%v (%v)",rojo,verde,blanco,data); time.Sleep(2*time.Second)
		}
	}
}
func menu5(){
	for b :=1; b>=1; b++{
		clear();data = " "
		crazytxt.Menu5()
		fmt.Printf("CrazyPop:");fmt.Scanf("%v",&data)
		if data == "salir"{
			salir()
		}else if data == "ayuda"{
			clear();crazytxt.Ayuda(); pausa()
		}else if data == "back"{
			break
		}else if data == "41"{
			crazyexec.Arat()
		}else if data == "42"{
			clear();proximamente();time.Sleep(2*time.Second)
		}else if data == "next"{
			clear();proximamente();time.Sleep(2*time.Second)
		}else{
			fmt.Printf("%vError en %v(¬!^-^)¬%v (%v)",rojo,verde,blanco,data); time.Sleep(2*time.Second)
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
func proximamente(){
	clear()
	fmt.Printf("%vProximamente%v...%v",amarillo,rojo,blanco)
	time.Sleep(2*time.Second)
}