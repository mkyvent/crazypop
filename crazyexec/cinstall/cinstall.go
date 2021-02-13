package cinstall

import "fmt"
import "os/exec"
import "time"
import "os"
var pause,tdata string
var (
	blanco string = "\033[37m"
	negro string = "\033[30m"
	fnegro string = "\033[40m"
	famarillo string = "\033[43m"
	fverde string = "\033[42m"
)
var (
	rojo string = "\033[31m"
	verde string = "\033[32m"
	amarillo string = "\033[33m"
	azul string = "\033[36m"
)
//########### Instalacion de las aplicaciones a termux
func Nmap(){
	for k := 1; k >=1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/nmap"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			nmp := exec.Command("apt","install","nmap","-y"); nmp.Stdout = os.Stdout; nmp.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Yow(){
	for k := 1; k >=1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/YOW"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando Yow...\n")
					yoww := exec.Command("git","clone","https://github.com/juanvelascogomez/YOW.git"); yoww.Stdout = os.Stdout; yoww.Run()
					mvy := exec.Command("mv","YOW","/data/data/com.termux/files/home/YOW"); mvy.Stdout = os.Stdout; mvy.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Git...\n")
				gtd := exec.Command("apt","install","git","-y"); gtd.Stdout = os.Stdout; gtd.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			pydos := exec.Command("apt","install","python2","-y"); pydos.Stdout = os.Stdout; pydos.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Youtubedl(){
	for k := 1; k >=1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/ffmpeg"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/curl"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/usr/bin/youtube-dl"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando Youtube-dl...\n")
					cyl := exec.Command("curl","-L","https://yt-dl.org/downloads/latest/youtube-dl","-o","/data/data/com.termux/files/usr/bin/youtube-dl"); cyl.Stdout = os.Stdout; cyl.Run()
					chm := exec.Command("chmod","a+rx","/data/data/com.termux/files/usr/bin/youtube-dl"); chm.Stdout = os.Stdout; chm.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando curl...\n")
				curll := exec.Command("apt","install","curl","-y"); curll.Stdout = os.Stdout; curll.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			ffmt := exec.Command("apt","install","ffmpeg","-y"); ffmt.Stdout = os.Stdout; ffmt.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Moonbuggy(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/moon-buggy"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			bgg := exec.Command("apt","install","moon-buggy","-y"); bgg.Stdout = os.Stdout; bgg.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Clamav(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/clamscan"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			clmav := exec.Command("apt","install","clamav","-y"); clmav.Stdout = os.Stdout; clmav.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Hydra(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/hydra"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			hy := exec.Command("apt","install","hydra","-y"); hy.Stdout = os.Stdout; hy.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Ubuntu(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/wget"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/openssl"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/usr/bin/proot"); err == nil{
					if _,err := os.Stat("/data/data/com.termux/files/home/ubuntu.sh"); err == nil{
						exitoso()
						time.Sleep(3*time.Second)
						break
					}else if os.IsNotExist(err){
						fmt.Printf("(^o^) Instalando Ubuntu\n")
						ssh := exec.Command("hash","-r"); ssh.Stdout = os.Stdout; ssh.Run()
						dbun := exec.Command("wget","https://raw.githubusercontent.com/EXALAB/AnLinux-Resources/master/Scripts/Installer/Ubuntu/ubuntu.sh"); dbun.Stdout = os.Stdout; dbun.Run()
						mvbun := exec.Command("mv","ubuntu.sh","/data/data/com.termux/files/home"); mvbun.Stdout = os.Stdout; mvbun.Run()
					}
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando proot...\n")
					prt := exec.Command("apt","install","proot","-y"); prt.Stdout = os.Stdout; prt.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando openssl-tool...\n")
				opn := exec.Command("apt","install","openssl-tool","-y"); opn.Stdout = os.Stdout; opn.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			wgt := exec.Command("apt","install","wget","-y"); wgt.Stdout = os.Stdout; wgt.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Kali(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/wget"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/openssl"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/usr/bin/proot"); err == nil{
					if _,err := os.Stat("/data/data/com.termux/files/home/kali.sh"); err == nil{
						exitoso()
						time.Sleep(3*time.Second)
						break
					}else if os.IsNotExist(err){
						fmt.Printf("(^o^) Instalando Kali...\n")
						hsss := exec.Command("hash","-r"); hsss.Stdout = os.Stdout; hsss.Run()
						dbu := exec.Command("wget","https://raw.githubusercontent.com/EXALAB/AnLinux-Resources/master/Scripts/Installer/Kali/kali.sh"); dbu.Stdout = os.Stdout; dbu.Run()
						mvkli := exec.Command("mv","kali.sh","/data/data/com.termux/files/home"); mvkli.Stdout = os.Stdout; mvkli.Run()
					}
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando proot...\n")
					prtt := exec.Command("apt","install","proot","-y"); prtt.Stdout = os.Stdout; prtt.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando openssl-tool...\n")
				opnn := exec.Command("apt","install","openssl-tool","-y"); opnn.Stdout = os.Stdout; opnn.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			wgtt := exec.Command("apt","install","wget","-y"); wgtt.Stdout = os.Stdout; wgtt.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Pacman(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/pacman"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			hyy := exec.Command("apt","install","pacman4console","-y"); hyy.Stdout = os.Stdout; hyy.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Freezer(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/wget"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/home/freeze.sh"); err == nil{
				exitoso()
				time.Sleep(3*time.Second)
				break
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Descargando Freeze.sh ...\n")
				fr := exec.Command("wget","https://github.com/Gameye98/V1RU5/raw/master/freeze.sh"); fr.Stdout = os.Stdout; fr.Run()
				mvf := exec.Command("mv","freeze.sh","/data/data/com.termux/files/home"); mvf.Stdout = os.Stdout; mvf.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			wggt := exec.Command("apt","install","wget","-y"); wggt.Stdout = os.Stdout; wggt.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Darksploit(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/home/DarkSploit"); err == nil{
				exitoso()
				time.Sleep(3*time.Second)
				break
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando DarkSploit...\n")
				dowdrk := exec.Command("git","clone","https://github.com/anthrax3/DarkSploit.git"); dowdrk.Stdout = os.Stdout; dowdrk.Run()
				mvdrk := exec.Command("mv","DarkSploit","/data/data/com.termux/files/home"); mvdrk.Stdout = os.Stdout; mvdrk.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			gitt := exec.Command("apt","install","git","-y"); gitt.Stdout = os.Stdout; gitt.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Tor(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/tor"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			tor := exec.Command("apt","install","tor","-y"); tor.Stdout = os.Stdout; tor.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Weeman(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/weeman"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando Weeman...\n")
					piph := exec.Command("python2","-m","pip","ps4","html5lib","lxml"); piph.Stdout = os.Stdout; piph.Run()
					weed := exec.Command("git","clone","https://github.com/evait-security/weeman"); weed.Stdout = os.Stdout; weed.Run()
					mvwee := exec.Command("mv","weeman","/data/data/com.termux/files/home"); mvwee.Stdout = os.Stdout; mvwee.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Python 2.7...\n")
				py2 := exec.Command("apt","install","python2","clang","-y"); py2.Stdout = os.Stdout; py2.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			wgit := exec.Command("apt","install","git","-y"); wgit.Stdout = os.Stdout; wgit.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Crunch(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/crunch"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			crun := exec.Command("apt","install","crunch","-y"); crun.Stdout = os.Stdout; crun.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Debian(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/wget"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/openssl"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/usr/bin/proot"); err == nil{
					if _,err := os.Stat("/data/data/com.termux/files/home/debian.sh"); err == nil{
						exitoso()
						time.Sleep(3*time.Second)
						break
					}else if os.IsNotExist(err){
						fmt.Printf("(^o^) Instalando Debian...\n")
						hsssd := exec.Command("hash","-r"); hsssd.Stdout = os.Stdout; hsssd.Run()
						db := exec.Command("wget","https://raw.githubusercontent.com/EXALAB/AnLinux-Resources/master/Scripts/Installer/Debian/debian.sh"); db.Stdout = os.Stdout; db.Run()
						mvklib := exec.Command("mv","debian.sh","/data/data/com.termux/files/home"); mvklib.Stdout = os.Stdout; mvklib.Run()
					}
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando proot...\n")
					prtl := exec.Command("apt","install","proot","-y"); prtl.Stdout = os.Stdout; prtl.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando openssl-tool...\n")
				opnnl := exec.Command("apt","install","openssl-tool","-y"); opnnl.Stdout = os.Stdout; opnnl.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			wgttp := exec.Command("apt","install","wget","-y"); wgttp.Stdout = os.Stdout; wgttp.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Exiftool(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/exiftool"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			xftl := exec.Command("apt","install","exiftool","-y"); xftl.Stdout = os.Stdout; xftl.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Gnupg(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/gpg"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			gpg := exec.Command("apt","install","gnupg","-y"); gpg.Stdout = os.Stdout; gpg.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Cupp(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/cupp"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando Cupp...\n")
					dcupp := exec.Command("git","clone","https://github.com/Mebus/cupp"); dcupp.Stdout = os.Stdout; dcupp.Run()
					mvcpp := exec.Command("mv","cupp","/data/data/com.termux/files/home"); mvcpp.Stdout = os.Stdout; mvcpp.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Python 2.7...\n")
				py3 := exec.Command("apt","install","python2","-y"); py3.Stdout = os.Stdout; py3.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			gti := exec.Command("apt","install","git","-y"); gti.Stdout = os.Stdout; gti.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Infoga(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/Infoga"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando Infoga...\n")
					pipnt := exec.Command("python2","-m","pip","install","requests","urllib3","urlparse"); pipnt.Stdout = os.Stdout; pipnt.Run()
					dwin := exec.Command("git","clone","https://github.com/m4ll0k/Infoga"); dwin.Stdout = os.Stdout; dwin.Run()
					mvin := exec.Command("mv","Infoga","/data/data/com.termux/files/home"); mvin.Stdout = os.Stdout; mvin.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Python 2.7...\n")
				pytw := exec.Command("apt","install","python2","-y"); pytw.Stdout = os.Stdout; pytw.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			gtk := exec.Command("apt","install","git","-y"); gtk.Stdout = os.Stdout; gtk.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Phoneinfoga(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/PhoneInfoga"); err == nil{
					exitoso()
					time.Sleep(1*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando PhoneInfoga...\n")
					pipg := exec.Command("pip","install","--upgrade","pip"); pipg.Stdout = os.Stdout; pipg.Run()
					dwphn := exec.Command("git","clone","https://github.com/ExpertAnonymous/PhoneInfoga"); dwphn.Stdout = os.Stdout; dwphn.Run()
					mvjk := exec.Command("mv","PhoneInfoga","/data/data/com.termux/files/home/"); mvjk.Stdout = os.Stdout; mvjk.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Python...\n")
				pyo := exec.Command("apt","install","python","-y"); pyo.Stdout = os.Stdout; pyo.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			gik := exec.Command("apt","install","git","-y"); gik.Stdout = os.Stdout; gik.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Tmux(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/tmux"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			tmx := exec.Command("apt","install","tmux","-y"); tmx.Stdout = os.Stdout; tmx.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Torshammer(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/torshammer"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando torshammer...\n")
					dwmer := exec.Command("git","clone","https://github.com/dotfighter/torshammer"); dwmer.Stdout = os.Stdout; dwmer.Run()
					mvtor := exec.Command("mv","torshammer","/data/data/com.termux/files/home"); mvtor.Stdout = os.Stdout; mvtor.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando python 2.7...\n")
				pyth2 := exec.Command("apt","install","python2","-y"); pyth2.Stdout = os.Stdout; pyth2.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			gjk := exec.Command("apt","install","git","-y"); gjk.Stdout = os.Stdout; gjk.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Slowloris(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/slowloris"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando Slowloris...\n")
					dws := exec.Command("git","clone","https://github.com/gkbrk/slowloris"); dws.Stdout = os.Stdout; dws.Run()
					mvsl := exec.Command("mv","slowloris","/data/data/com.termux/files/home"); mvsl.Stdout = os.Stdout; mvsl.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando python 2.7 ...\n")
				spy2 := exec.Command("apt","install","python2","-y"); spy2.Stdout = os.Stdout; spy2.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			sgit := exec.Command("apt","install","git","-y"); sgit.Stdout = os.Stdout; sgit.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Spiderbot(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/php"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/SpiderBot"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando SpiderBot...\n")
					dws := exec.Command("git","clone","https://github.com/Cvar1984/SpiderBot"); dws.Stdout = os.Stdout; dws.Run()
					mvsl := exec.Command("mv","SpiderBot","/data/data/com.termux/files/home"); mvsl.Stdout = os.Stdout; mvsl.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando PHP..\n")
				phpd := exec.Command("apt","install","php","-y"); phpd.Stdout = os.Stdout; phpd.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			spid := exec.Command("apt","install","git","-y"); spid.Stdout = os.Stdout; spid.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Sqlmap(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/sqlmap"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando SQLmap...\n")
					dws := exec.Command("git","clone","https://github.com/sqlmapproject/sqlmap"); dws.Stdout = os.Stdout; dws.Run()
					mvsl := exec.Command("mv","sqlmap","/data/data/com.termux/files/home"); mvsl.Stdout = os.Stdout; mvsl.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Python 2.7 ...\n")
				pythn := exec.Command("apt","install","python2","-y"); pythn.Stdout = os.Stdout; pythn.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			gitsl := exec.Command("apt","install","git","-y"); gitsl.Stdout = os.Stdout; gitsl.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Cmatrix(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/cmatrix"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			trx := exec.Command("apt","install","cmatrix","-y"); trx.Stdout = os.Stdout; trx.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func W3m(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/w3m"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			wm3 := exec.Command("apt","install","w3m","-y"); wm3.Stdout = os.Stdout; wm3.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Htop(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/htop"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			top := exec.Command("apt","install","htop","-y"); top.Stdout = os.Stdout; top.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Hashcat(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/hashcat"); err == nil{
			
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			stl := exec.Command("apt","install","unstable-repo","-y"); stl.Stdout = os.Stdout; stl.Run()
			cat := exec.Command("apt","install","hashcat","-y"); cat.Stdout = os.Stdout; cat.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Lazymux(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/Lazymux"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando Lazymux...\n")
					lzm := exec.Command("git","clone","https://github.com/Gameye98/Lazymux"); lzm.Stdout = os.Stdout; lzm.Run()
					mvlzm := exec.Command("mv","Lazymux","/data/data/com.termux/files/home"); mvlzm.Stdout = os.Stdout; mvlzm.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Git...\n")
				gkh := exec.Command("apt","install","git","-y"); gkh.Stdout = os.Stdout; gkh.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			thm := exec.Command("apt","install","python","-y"); thm.Stdout = os.Stdout; thm.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Wpspin(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/WPSpin"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando WPSpin...\n")
					lzmm := exec.Command("git","clone","https://github.com/drygdryg/WPSpin"); lzmm.Stdout = os.Stdout; lzmm.Run()
					mvlzmn := exec.Command("mv","WPSpin","/data/data/com.termux/files/home"); mvlzmn.Stdout = os.Stdout; mvlzmn.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Git...\n")
				gkhh := exec.Command("apt","install","git","-y"); gkhh.Stdout = os.Stdout; gkhh.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			pyh := exec.Command("apt","install","python","-y"); pyh.Stdout = os.Stdout; pyh.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Commix(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/commix"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando Commix...\n")
					lzmmk := exec.Command("git","clone","https://github.com/commixproject/commix"); lzmmk.Stdout = os.Stdout; lzmmk.Run()
					mvlzmnk := exec.Command("mv","commix","/data/data/com.termux/files/home"); mvlzmnk.Stdout = os.Stdout; mvlzmnk.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Git...\n")
				gkkh := exec.Command("apt","install","git","-y"); gkkh.Stdout = os.Stdout; gkkh.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			pyhh := exec.Command("apt","install","python2","-y"); pyhh.Stdout = os.Stdout; pyhh.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Sms(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/SpazSMS"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando SpazSMS...\n")
					pip2 := exec.Command("pip2","install","--upgrade","pip"); pip2.Stdout = os.Stdout; pip2.Run()
					pir := exec.Command("python2","-m","pip","install","requests"); pir.Stdout = os.Stdout; pir.Run()
					lzmmk := exec.Command("git","clone","https://github.com/Gameye98/SpazSMS"); lzmmk.Stdout = os.Stdout; lzmmk.Run()
					mvlzmnk := exec.Command("mv","SpazSMS","/data/data/com.termux/files/home"); mvlzmnk.Stdout = os.Stdout; mvlzmnk.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Git...\n")
				ggkkh := exec.Command("apt","install","git","-y"); ggkkh.Stdout = os.Stdout; ggkkh.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			pylhh := exec.Command("apt","install","python2","-y"); pylhh.Stdout = os.Stdout; pylhh.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Neofetch(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/neofetch"); err == nil{
			exitoso()
			time.Sleep(3*time.Second)
			break
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			neo := exec.Command("apt","install","neofetch","-y"); neo.Stdout = os.Stdout; neo.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Flappy(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/flappy_bird"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando Flappy-Bird...\n")
					yuj := exec.Command("git","clone","https://github.com/JustAHackers/flappy_bird"); yuj.Stdout = os.Stdout; yuj.Run()
					mvfp := exec.Command("mv","flappy_bird","/data/data/com.termux/files/home/"); mvfp.Stdout = os.Stdout; mvfp.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Python 2.7 ...\n")
				ght := exec.Command("apt","install","python2","-y"); ght.Stdout = os.Stdout; ght.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			neoo := exec.Command("apt","install","git","-y"); neoo.Stdout = os.Stdout; neoo.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Vbug(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/vbug"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando Vbug...\n")
					yujj := exec.Command("git","clone","https://github.com/Gameye98/vbug"); yujj.Stdout = os.Stdout; yujj.Run()
					mvfpp := exec.Command("mv","vbug","/data/data/com.termux/files/home/"); mvfpp.Stdout = os.Stdout; mvfpp.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Python 2.7...\n")
				gght := exec.Command("apt","install","python2","-y"); gght.Stdout = os.Stdout; gght.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			neooo := exec.Command("apt","install","git","-y"); neooo.Stdout = os.Stdout; neooo.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Routersploit(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/routersploit"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando RouterSploit...\n")
					yujjk := exec.Command("git","clone","https://github.com/threat9/routersploit"); yujjk.Stdout = os.Stdout; yujjk.Run()
					mvfpk := exec.Command("mv","routersploit","/data/data/com.termux/files/home/"); mvfpk.Stdout = os.Stdout; mvfpk.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Python 2.7 ...\n")
				gghtk := exec.Command("apt","install","python2","-y"); gghtk.Stdout = os.Stdout; gghtk.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			neool := exec.Command("apt","install","git","-y"); neool.Stdout = os.Stdout; neool.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Blackhydra(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/usr/bin/hydra"); err == nil{
					if _,err := os.Stat("/data/data/com.termux/files/home/Black-Hydra"); err == nil{
						exitoso()
						time.Sleep(3*time.Second)
						break
					}else if os.IsNotExist(err){
						fmt.Printf("(^o^) Instalando Black-Hydra...\n")
						akl := exec.Command("git","clone","https://github.com/Gameye98/Black-Hydra"); akl.Stdout = os.Stdout; akl.Run()
						mvhyd := exec.Command("mv","Black-Hydra","/data/data/com.termux/files/home/"); mvhyd.Stdout = os.Stdout; mvhyd.Run()
					}
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando Hydra...\n")
					sds := exec.Command("apt","install","hydra","-y"); sds.Stdout = os.Stdout; sds.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Python 2.7...\n")
				kkp := exec.Command("apt","install","python2","-y"); kkp.Stdout = os.Stdout; kkp.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			gtm := exec.Command("apt","install","git","-y"); gtm.Stdout = os.Stdout; gtm.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Websploit(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/home/websploit"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando WebSploit ...\n")
					pipdw := exec.Command("python2","-m","pip","install","scapy"); pipdw.Stdout = os.Stdout; pipdw.Run()
					dwnr := exec.Command("git","clone"," https://github.com/The404Hacking/websploit"); dwnr.Stdout = os.Stdout; dwnr.Run()
					mbweb := exec.Command("mv","websploit","/data/data/com.termux/files/home/"); mbweb.Stdout = os.Stdout; mbweb.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando Python 2.7 ...\n")
				pydos := exec.Command("apt","install","python2","-y"); pydos.Stdout = os.Stdout; pydos.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			gtwz := exec.Command("apt","install","git","-y"); gtwz.Stdout = os.Stdout; gtwz.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Astranmap(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/home/AstraNmap"); err == nil{
				exitoso()
				time.Sleep(3*time.Second)
				break
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando AstraNmap...\n")
				dwastra := exec.Command("git","clone","https://github.com/Gameye98/AstraNmap"); dwastra.Stdout = os.Stdout; dwastra.Run()
				mvastra := exec.Command("mv","AstraNmap","/data/data/com.termux/files/home"); mvastra.Stdout = os.Stdout; mvastra.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			gastra := exec.Command("apt","install","git","-y"); gastra.Stdout = os.Stdout; gastra.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
		}
	}
}
func Arat(){
	for k := 1; k >= 1; k++{
		clear()
		if _,err := os.Stat("/data/data/com.termux/files/usr/bin/git"); err == nil{
			if _,err := os.Stat("/data/data/com.termux/files/home/A-Rat"); err == nil{
				if _,err := os.Stat("/data/data/com.termux/files/usr/bin/python2"); err == nil{
					exitoso()
					time.Sleep(3*time.Second)
					break
				}else if os.IsNotExist(err){
					fmt.Printf("(^o^) Instalando Python 2.7 ...\n")
					py2t := exec.Command("apt","install","python2","-y"); py2t.Stdout = os.Stdout; py2t.Run()
				}
			}else if os.IsNotExist(err){
				fmt.Printf("(^o^) Instalando A-Rat...\n")
				dwastra := exec.Command("git","clone","https://github.com/RexTheGod/A-Rat.git"); dwastra.Stdout = os.Stdout; dwastra.Run()
				mvastra := exec.Command("mv","A-Rat","/data/data/com.termux/files/home"); mvastra.Stdout = os.Stdout; mvastra.Run()
			}
		}else if os.IsNotExist(err){
			instalando(); time.Sleep(1*time.Second)
			gastrra := exec.Command("apt","install","git","-y"); gastrra.Stdout = os.Stdout; gastrra.Run()
		}
		if k >= 20{
			clear()
			fmt.Printf("%vError en instalacion...   Regresando...%v",rojo,blanco)
			time.Sleep(3*time.Second)
			break
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
func exitoso(){
	clear()
	fmt.Printf("%v _________________________\n",blanco)
	fmt.Printf("%v|%v%v_CrazyPop_______[_][-][X]%v%v|\n",blanco,negro,fverde,fnegro,blanco)
	fmt.Printf("%v|%v%v                         %v%v|\n",blanco,negro,fverde,fnegro,blanco)
	fmt.Printf("%v|%v%v       Exitosa !!!       %v%v|\n",blanco,negro,fverde,fnegro,blanco)
	fmt.Printf("%v|%v%v_________________________%v%v|\n\n",blanco,negro,fverde,fnegro,blanco)
	fmt.Printf("Instalado en el sistema ...\n")
}
func instalando(){
	clear()
	fmt.Printf("%v _________________________\n",blanco)
	fmt.Printf("%v|%v%v_CrazyPop_______[_][-][X]%v%v|\n",blanco,negro,famarillo,fnegro,blanco)
	fmt.Printf("%v|%v%v                         %v%v|\n",blanco,negro,famarillo,fnegro,blanco)
	fmt.Printf("%v|%v%v       Instalando...     %v%v|\n",blanco,negro,famarillo,fnegro,blanco)
	fmt.Printf("%v|%v%v_________________________%v%v|\n\n",blanco,negro,famarillo,fnegro,blanco)
}