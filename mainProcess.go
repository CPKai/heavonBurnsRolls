package main

import (
	"fmt"
	"os"

	"github.com/go-vgo/robotgo"
)

var (
	screenWidth     int     = 1920
	screenLength    int     = 1080
	tolerance       float64 = 0.2
	chooseLang      string  = "img/step-chooseLang.png"
	autoCtrl        string  = "img/step-autoCtrl.png"
	nextPng         string  = "img/step-next.png"
	firstGacha      string  = "img/step-firstGacha.png"
	firstGacha2     string  = "img/step-firstGacha2.png"
	teachHand       string  = "img/step-teachHand.png"
	storyChoose     string  = "img/step-storyChoose.png"
	atGachaPage     string  = "img/step-atGachaPage.png"
	abPool          string  = "img/gacha-abPool.png"
	abPool2         string  = "img/gacha-abPool2.png"
	rollAgain       string  = "img/gacha-rollAgain.png"
	pinkPool        string  = "img/gacha-pinkPool.png"
	bluePool        string  = "img/gacha-bluePool.png"
	ok_png          string  = "img/step-ok.png"
	mainPage        string  = "img/step-mainPage.png"
	target_role_1   string  = "img/role-angel.png"
	team            string  = "img/step-team.png"
	teamFilter      string  = "img/step-teamFilter.png"
	teamFilterSS    string  = "img/step-teamFilterSS.png"
	settings        string  = "img/step-settings.png"
	menuOfFirstPage string  = "img/step-menuOfFirstPage.png"
	deleteAccount   string  = "img/step-deleteAccount.png"
	deleteAccount2  string  = "img/step-deleteAccount2.png"
)

func main() {
	fmt.Println("Heaven Burns Red 工具箱")
	fmt.Println("按f1停止程式")
	fmt.Println("開始進行首抽刷取")
	go exitEvent()
	rollGacha()
}

func rollGacha() {
	var (
		// cur_phase     string = ""
		round_counter int = 0
		// PHASE_BEGIN   int    = 0
		// PHASE_STORY   int    = 1
		// PHASE_GACHA   int    = 2
		// PHASE_CHECK   int    = 3
		// PHASE_DEL     int    = 4
		// target_role_1 string = "img/role-angel.png"
	)
	for true {
		round_counter++
		fmt.Printf("current round:%d\n", round_counter)
		phaseBegin()
		phaseStory()
		phaseGacha()
		result := phaseCheck()
		if result {
			fmt.Println("已抽到天使")
			break
		}
		phaseDel()
	}

}

func findImageOnScreen(phaseNum float64, imgPath string, sp_tolerance float64, useESC bool, useEnter bool, clickTarget bool) {
	fmt.Printf("phaseNum:%v\n", phaseNum)
	fmt.Printf("imgPath:%s\n", imgPath)
	for true {
		if useESC {
			robotgo.KeyTap("esc")
			robotgo.Sleep(1)
		}
		if useEnter {
			robotgo.KeyTap("enter")
			robotgo.Sleep(1)
		}
		robotgo.Sleep(1)
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(imgPath, bitmap_screen, sp_tolerance)
		if fx != -1 && fy != -1 {
			if clickTarget {
				robotgo.MoveClick(fx, fy, "left", false)
			}
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
}

func phaseBegin() {
	bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
	robotgo.FreeBitmap(bitmap_screen)

	// wait 繁體中文 and click
	findImageOnScreen(1.1, chooseLang, tolerance, false, false, true)

	// ok
	// ok
	// 同意
	// clickToLogin(atFirstPage.png)
	findImageOnScreen(1.2, autoCtrl, tolerance, false, true, false)
}

func phaseStory() {
	var ()

	bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
	robotgo.FreeBitmap(bitmap_screen)

	// step - autoCtrl
	fmt.Println("p2.1")
	for true {
		robotgo.Sleep(1)
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(autoCtrl, bitmap_screen, tolerance-0.1)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// step - autoCtrl
	fmt.Println("p2.2")
	for true {
		robotgo.Sleep(1)
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(autoCtrl, bitmap_screen, tolerance-0.1)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// step - next
	// step - next
	// step - endTeach
	// step - tryGacha
	fmt.Println("p2.3")
	for true {
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(nextPng, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.KeyTap("enter")
			robotgo.Sleep(1)
			robotgo.KeyTap("enter")
			robotgo.Sleep(1)
			robotgo.KeyTap("esc")
			robotgo.Sleep(1)
			robotgo.KeyTap("enter")
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}

	// step - firstGacha
	fmt.Println("p2.4")
	for true {
		robotgo.Sleep(2)
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(firstGacha, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}

	// stop - firstGacha2
	fmt.Println("p2.5")
	for true {
		robotgo.Sleep(2)
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(firstGacha2, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// click loop until autoCtrl
	fmt.Println("p2.6")
	for true {
		robotgo.KeyTap("enter")
		robotgo.Sleep(1)
		robotgo.KeyTap("esc")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(autoCtrl, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
		robotgo.Sleep(1)
	}

	// wait until step-teachHand
	// teachHand
	// teachHand
	// teachHand
	// click one time
	// teachHand
	// teachHand(至此每個teachHand位置皆不同)
	fmt.Println("p2.7")
	teachHandCount := 0
	for true {
		bitmap_screen2 := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(teachHand, bitmap_screen2, tolerance)
		if fx != -1 && fy != -1 {
			fmt.Println("p2.71")
			teachHandCount = teachHandCount + 1
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen2)
			fmt.Println("p2.72")
			if teachHandCount == 4 {
				fmt.Println("p2.73")
				robotgo.KeyTap("enter")
			} else if teachHandCount == 6 {
				fmt.Println("p2.74")
				break
			}
		}
		fmt.Println("p2.75")
		robotgo.FreeBitmap(bitmap_screen2)
	}

	// click until autoCtrl
	fmt.Println("p2.8")
	for true {
		robotgo.KeyTap("esc")
		robotgo.KeyTap("enter")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(autoCtrl, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
		robotgo.Sleep(1)
	}

	// wait storyChoose
	fmt.Println("p2.9")
	for true {
		robotgo.KeyTap("esc")
		robotgo.KeyTap("enter")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(storyChoose, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
		robotgo.Sleep(1)
	}

	// wait autoCtrl
	fmt.Println("p2.10")
	for true {
		robotgo.KeyTap("esc")
		robotgo.KeyTap("enter")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(autoCtrl, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
		robotgo.Sleep(1)
	}

	// wait day00end
	// keep click until teachHand
	// teachHand
	// teachHand
	// teachHand
	// teachHand
	// teachHand
	fmt.Println("p2.11")
	teachHandCount = 0
	for true {
		robotgo.KeyTap("esc")
		robotgo.KeyTap("enter")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(teachHand, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			teachHandCount++
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			if teachHandCount == 6 {
				break
			}
		}
		robotgo.FreeBitmap(bitmap_screen)
		robotgo.Sleep(1)
	}
	// wait until atGachaPage

}

func phaseGacha() {
	var ()

GachaPageCheck:
	fmt.Println("p3.1")
	bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
	fx, fy := robotgo.FindPic(atGachaPage, bitmap_screen, tolerance)
	if fx != -1 && fy != -1 {
		robotgo.MoveClick(fx, fy, "left", false)
		robotgo.FreeBitmap(bitmap_screen)
		goto RollAngelBeatPool
	}
	robotgo.FreeBitmap(bitmap_screen)
	goto GachaPageCheck

RollAngelBeatPool:
	fmt.Println("p3.2")
	// E 3 times
	for i := 0; i < 3; i++ {
		robotgo.KeyTap("e")
		robotgo.Sleep(1)
	}

	// gacha-abPool
	for true {
		fmt.Println("p3.3")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(abPool, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// gacha-abPool2
	for true {
		fmt.Println("p3.4")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(abPool2, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// keep enter until rollAgain and click rollAgain
	for true {
		fmt.Println("p3.5")
		robotgo.KeyTap("enter")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(rollAgain, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// gacha-abPool2
	for true {
		fmt.Println("p3.6")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(abPool2, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// keep enter until rollAgain
	// esc
	for true {
		fmt.Println("p3.7")
		robotgo.KeyTap("enter")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(rollAgain, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.KeyTap("esc")
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// wait until gacha-abPool
	// E 3 times
	for true {
		fmt.Println("p3.8")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(abPool, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			for i := 0; i < 3; i++ {
				robotgo.KeyTap("e")
				robotgo.Sleep(1)
			}
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}

	// gacha-pinkPool
	for true {
		fmt.Println("p3.9")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(pinkPool, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// gacha-abPool2
	for true {
		fmt.Println("p3.10")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(abPool2, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// keep enter until ok and click it
	for true {
		fmt.Println("p3.11")
		robotgo.KeyTap("enter")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(ok_png, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// wait until atGachaPage
	// E 5 times
	for true {
		fmt.Println("p3.12")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(atGachaPage, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			for i := 0; i < 5; i++ {
				robotgo.KeyTap("e")
				robotgo.Sleep(1)
			}
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// gacha-bluePool
	for true {
		fmt.Println("p3.13")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(bluePool, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// gacha-abPool2
	for true {
		fmt.Println("p3.14")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(abPool2, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// keep enter until rollAgain and click rollAgain
	for true {
		fmt.Println("p3.15")
		robotgo.KeyTap("enter")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(rollAgain, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// gacha-abPool2
	for true {
		fmt.Println("p3.16")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(abPool2, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// keep enter until ok
	for true {
		fmt.Println("p3.17")
		robotgo.KeyTap("enter")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(ok_png, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}

	// wait until atGachaPage
	// esc
	for true {
		fmt.Println("p3.18")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(atGachaPage, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.KeyTap("esc")
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}

	// wait until mainPage
	// click mainPage (enter?)
	// enter(?)
	// keep esc until mainPage
	for true {
		fmt.Println("p3.19")
		robotgo.KeyTap("enter")
		robotgo.KeyTap("esc")
		bitmap_screen = robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy = robotgo.FindPic(mainPage, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
}

func phaseCheck() bool {
	var ()
	// 進選單
	for true {
		fmt.Println("p4.1")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(mainPage, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// 進小隊頁
	for true {
		fmt.Println("p4.2")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(team, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	//進篩選頁
	for true {
		fmt.Println("p4.3")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(teamFilter, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	//篩ss & click ok
	for true {
		fmt.Println("p4.4")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(teamFilterSS, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			break
		}
		fx, fy = robotgo.FindPic(ok_png, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
CheckHaveTargetOrNot:
	fmt.Println("p4.5")
	checkCounter := 0
	bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
	fx, fy := robotgo.FindPic(target_role_1, bitmap_screen, tolerance)
	if fx != -1 && fy != -1 {
		robotgo.FreeBitmap(bitmap_screen)
		return true
	}
	robotgo.FreeBitmap(bitmap_screen)
	if checkCounter < 3 {
		goto CheckHaveTargetOrNot
	}

	return false
}

func phaseDel() {
	var ()

	// click mainPage
	for true {
		fmt.Println("p5.1")
		robotgo.KeyTap("esc")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(mainPage, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// wait until settings and click settings
	for true {
		fmt.Println("p5.2")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(settings, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// keep enter until teachHand and click teachHand
	// wait until teachHand
	// wait until keyboardCheck
	// click enter
	// esc
	// 上
	// 上
	// enter
	// OK(or enter)
	teachHandCount := 0
	for true {
		fmt.Println("p5.3")
		robotgo.KeyTap("enter")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(teachHand, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			teachHandCount++
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			if teachHandCount == 2 {
				robotgo.KeyTap("enter")
				robotgo.Sleep(1)
				robotgo.KeyTap("esc")
				robotgo.Sleep(1)
				robotgo.KeyTap("up")
				robotgo.Sleep(1)
				robotgo.KeyTap("up")
				robotgo.Sleep(1)
				robotgo.KeyTap("enter")
				robotgo.Sleep(1)
				robotgo.KeyTap("enter")
				break
			}
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// wait menuOfFirstPage and click it
	for true {
		fmt.Println("p5.4")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(menuOfFirstPage, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// wait deleteAccount and click it
	for true {
		fmt.Println("p5.5")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(deleteAccount, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// wait deleteAccount2 and click it
	for true {
		fmt.Println("p5.6")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(deleteAccount2, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// OK
	for true {
		fmt.Println("p5.7")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(ok_png, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	// OK
	for true {
		fmt.Println("p5.8")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(ok_png, bitmap_screen, tolerance)
		if fx != -1 && fy != -1 {
			robotgo.MoveClick(fx, fy, "left", false)
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}

	// (接續phaseBegin)

}

func exitEvent() {
	exit_event := robotgo.AddEvent("f1")
	if exit_event {
		fmt.Println("終止程式")
		os.Exit(0)
	}
}

func errHandler(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}
