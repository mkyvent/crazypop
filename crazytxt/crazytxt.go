package crazytxt

import "fmt"
var (
	blanco string = "\033[37m"
	negro string = "\033[30m"
	verde string = "\033[32m"
	amarillo string = "\033[33m"
	azul string = "\033[36m"
	fazul string = "\033[46m"
	fnegro string = "\033[40m"
)
func Menu(){
	fmt.Printf("               %v%vC R A Z Y P O P%v%v         [%v1%v]\n",fazul,negro,fnegro,blanco,verde,blanco)
	fmt.Printf("   [%v1%v] Nmap (Scanner)      [%v6%v] THC Hydra\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v2%v] Yow                 [%v7%v] Ubuntu\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v3%v] youtube-dl          [%v8%v] kali\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v4%v] Moon-Buggy(Game)    [%v9%v] pacman (Game)\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v5%v] ClamAV(antimalware) [%v10%v]freezer(malware)\n",verde,blanco,verde,blanco)
	fmt.Printf("       [%vsalir%v]        [%vayuda%v]       [%vnext%v]\n",amarillo,blanco,amarillo,blanco,amarillo,blanco)
}
func Menu2(){
	fmt.Printf("               %v%vC R A Z Y P O P%v%v         [%v2%v]\n",fazul,negro,fnegro,blanco,verde,blanco)
	fmt.Printf("   [%v11%v] Darksploit          [%v16%v] Exfitool\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v12%v] TOR                 [%v17%v] Gnupg\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v13%v] Weeman              [%v18%v] Cupp\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v14%v] Crunch              [%v19%v] Infoga\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v15%v] Debian              [%v20%v]PhoneInfoga\n",verde,blanco,verde,blanco)
	fmt.Printf("  [%vback%v]     [%vsalir%v]       [%vayuda%v]      [%vnext%v]\n",amarillo,blanco,amarillo,blanco,amarillo,blanco,amarillo,blanco)
}
func Menu3(){
	fmt.Printf("               %v%vC R A Z Y P O P%v%v         [%v3%v]\n",fazul,negro,fnegro,blanco,verde,blanco)
	fmt.Printf("   [%v21%v] Tmux                [%v26%v] Cmatrix\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v22%v] Torshammer          [%v27%v] W3m\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v23%v] Slowloris           [%v28%v] htop\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v24%v] SpiderBot           [%v29%v] Hashcat\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v25%v] SQLmap              [%v30%v] Lazymux\n",verde,blanco,verde,blanco)
	fmt.Printf("  [%vback%v]     [%vsalir%v]       [%vayuda%v]      [%vnext%v]\n",amarillo,blanco,amarillo,blanco,amarillo,blanco,amarillo,blanco)
}
func Menu4(){
	fmt.Printf("               %v%vC R A Z Y P O P%v%v         [%v4%v]\n",fazul,negro,fnegro,blanco,verde,blanco)
	fmt.Printf("   [%v31%v] Wpspin              [%v36%v] Vbug\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v32%v] Commix              [%v37%v] RouterSploit\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v33%v] SpazSMS             [%v38%v] Black_Hydra\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v34%v] Neofetch            [%v39%v] WebSploit\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v35%v] Flappy_Bird(game)   [%v40%v] astraNmap\n",verde,blanco,verde,blanco)
	fmt.Printf("  [%vback%v]     [%vsalir%v]       [%vayuda%v]      [%vnext%v]\n",amarillo,blanco,amarillo,blanco,amarillo,blanco,amarillo,blanco)
}
func Menu5(){
	fmt.Printf("               %v%vC R A Z Y P O P%v%v         [%v5%v]\n",fazul,negro,fnegro,blanco,verde,blanco)
	fmt.Printf("   [%v41%v] A-Rat               [%v46%v] *\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v42%v] *                   [%v47%v] *\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v43%v] *                   [%v48%v] *\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v44%v] *                   [%v49%v] *\n",verde,blanco,verde,blanco)
	fmt.Printf("   [%v45%v] *                   [%v50%v] *\n",verde,blanco,verde,blanco)
	fmt.Printf("  [%vback%v]     [%vsalir%v]       [%vayuda%v]      [%vnext%v]\n",amarillo,blanco,amarillo,blanco,amarillo,blanco,amarillo,blanco)
}
func Logo(){
	fmt.Printf("%v ___  ___  ___  ___ ___ ___\n",azul)
	fmt.Printf("|   || __||   ||_  ||  V  |\n")
	fmt.Printf("| --|| |  | - ||  _| -   -\n")
	fmt.Printf("|___||_|  |_|_||___|  |_|\n")
	fmt.Printf("%v      ____    ___   ____\n",verde)
	fmt.Printf("     |     - -   - |    -\n")
	fmt.Printf("     |   0 ||     ||   0 |\n")
	fmt.Printf("     |  _-  |  0  ||  _-\n")
	fmt.Printf("     |_|     -___- |_|%v\n",blanco)
}
func Ayuda(){
	Logo()
	fmt.Printf("bienvenido a crazypop\n")
	fmt.Printf("es un software que ayuda a instalar\n")
	fmt.Printf("todo tipo de herramientas para\n")
	fmt.Printf("termux.\n")
	fmt.Printf("todo archivo ejecutable se instala en /usr/bin/\n")
	fmt.Printf("y los archivo en carpeta se instalan en /home/\n")
	fmt.Printf("el metodo de instalacion es facil, comprueba\n")
	fmt.Printf("la existencia del archivo primero y\n")
	fmt.Printf("al recibir una respuesta negativa,\n")
	fmt.Printf("crazypop procede a instalar el archivo\n")
	fmt.Printf("faltante, asi sucesivamente en cada\n")
	fmt.Printf("dependecia del archivo que necesite\n")
	fmt.Printf("y asi crazypop instala mejor cada \n")
	fmt.Printf("archivo que ocupes.\n")
	fmt.Printf("Creado por Miguel Ventura\n")
	
}
