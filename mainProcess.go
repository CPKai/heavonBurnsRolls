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
	teachHand1      string  = "img/step-teachHand1.png"
	teachHand2      string  = "img/step-teachHand2.png"
	teachHand3      string  = "img/step-teachHand3.png"
	teachHand4      string  = "img/step-teachHand4.png"
	teachHand5      string  = "img/step-teachHand5.png"
	teachHand6      string  = "img/step-teachHand6.png"
	teachHand7      string  = "img/step-teachHand7.png"
	teachHand8      string  = "img/step-teachHand8.png"
	teachHand9      string  = "img/step-teachHand9.png"
	teachHand10     string  = "img/step-teachHand10.png"
	teachHand11     string  = "img/step-teachHand11.png"
	teachHand12     string  = "img/step-teachHand12.png"
	teachHand13     string  = "img/step-teachHand13.png"
	teachHand14     string  = "img/step-teachHand14.png"
	returnFirstPage string  = "img/step-returnFirstPage.png"
	gachaResult     string  = "img/gacha-result.png"
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

func findImageOnScreen(phaseNum string, imgPath string, sp_tolerance float64, useESC bool, useEnter bool, clickTarget bool) {
	fmt.Printf("phaseNum:%v\t", phaseNum)
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
		robotgo.KeyTap("lctrl")
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
		// robotgo.KeyUp("lctrl")
	}
}

func phaseBegin() {
	// wait 繁體中文 and click
	findImageOnScreen("1.1", chooseLang, tolerance, false, false, true)

	// ok
	// ok
	// 同意
	// clickToLogin(atFirstPage.png)
	findImageOnScreen("1.2", autoCtrl, tolerance, false, true, false)
}

func phaseStory() {

	// step - autoCtrl
	findImageOnScreen("2.1", autoCtrl, tolerance-0.1, false, false, true)
	// step - autoCtrl
	findImageOnScreen("2.2", autoCtrl, tolerance-0.1, false, false, true)
	// step - next
	// step - next
	// step - endTeach
	// step - tryGacha

	findImageOnScreen("2.3", nextPng, tolerance, false, false, true)
	findImageOnScreen("2.4", nextPng, tolerance, false, false, true)
	robotgo.Sleep(1)
	robotgo.KeyTap("esc")
	robotgo.Sleep(1)
	robotgo.KeyTap("enter")

	// fmt.Println("p2.3")
	// for true {
	// 	bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
	// 	fx, fy := robotgo.FindPic(nextPng, bitmap_screen, tolerance)
	// 	if fx != -1 && fy != -1 {
	// 		robotgo.MoveClick(fx, fy, "left", false)
	// 		robotgo.KeyTap("enter")
	// 		robotgo.Sleep(1)
	// 		robotgo.KeyTap("enter")
	// 		robotgo.Sleep(1)
	// 		robotgo.KeyTap("esc")
	// 		robotgo.Sleep(1)
	// 		robotgo.KeyTap("enter")
	// 		robotgo.FreeBitmap(bitmap_screen)
	// 		break
	// 	}
	// 	robotgo.FreeBitmap(bitmap_screen)
	// }

	// step - firstGacha
	findImageOnScreen("2.5", firstGacha, tolerance, false, false, true)

	// stop - firstGacha2
	robotgo.Sleep(2)
	findImageOnScreen("2.6", firstGacha2, tolerance, false, false, true)
	robotgo.Sleep(2)

	// findImageOnScreen(2.51, ok_png, tolerance+0.05, true, true, true)
	// findImageOnScreen(2.6, autoCtrl, tolerance-0.1, true, true, true)

	// click loop until autoCtrl
	// findImageOnScreen(2.61, autoCtrl, tolerance-0.1, true, true, true)

	// wait until step-teachHand
	// teachHand
	// teachHand
	// teachHand
	// click one time
	// teachHand
	// teachHand(至此每個teachHand位置皆不同)
	findImageOnScreen("2.7", teachHand1, tolerance, true, true, true)
	findImageOnScreen("2.8", teachHand2, tolerance, false, false, true)
	findImageOnScreen("2.9", teachHand3, tolerance, false, false, true)
	findImageOnScreen("2.10", teachHand4, tolerance, false, false, true)
	robotgo.Sleep(2)
	robotgo.KeyTap("enter")
	findImageOnScreen("2.11", teachHand5, tolerance, false, false, true)
	findImageOnScreen("2.12", teachHand6, tolerance, false, false, true)

	// click until autoCtrl
	findImageOnScreen("2.13", autoCtrl, tolerance, true, true, true)

	// wait storyChoose
	findImageOnScreen("2.14", storyChoose, tolerance, false, false, true)

	// wait autoCtrl
	findImageOnScreen("2.15", autoCtrl, tolerance, true, true, true)

	// wait day00end
	// keep click until teachHand
	// teachHand
	// teachHand
	// teachHand
	// teachHand
	// teachHand
	findImageOnScreen("2.16", teachHand7, tolerance, false, true, true)
	findImageOnScreen("2.17", teachHand8, tolerance, false, false, true)
	findImageOnScreen("2.18", teachHand9, tolerance, false, false, true)
	robotgo.Sleep(2)
	findImageOnScreen("2.19", teachHand10, tolerance, false, false, true)
	robotgo.Sleep(2)
	findImageOnScreen("2.20", teachHand11, tolerance, false, false, true)
	findImageOnScreen("2.21", teachHand12, tolerance, false, false, true)
	// fmt.Println("p2.11")
	// teachHandCount = 0
	// for true {
	// 	robotgo.KeyTap("esc")
	// 	robotgo.KeyTap("enter")
	// 	bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
	// 	fx, fy := robotgo.FindPic(teachHand, bitmap_screen, tolerance)
	// 	if fx != -1 && fy != -1 {
	// 		teachHandCount++
	// 		robotgo.MoveClick(fx, fy, "left", false)
	// 		robotgo.FreeBitmap(bitmap_screen)
	// 		if teachHandCount == 6 {
	// 			break
	// 		}
	// 	}
	// 	robotgo.FreeBitmap(bitmap_screen)
	// 	robotgo.Sleep(1)
	// }
	// wait until atGachaPage

}

func phaseGacha() {
	findImageOnScreen("3.1", atGachaPage, tolerance, false, false, false)
	// GachaPageCheck:
	// 	fmt.Println("p3.1")
	// 	bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
	// 	fx, fy := robotgo.FindPic(atGachaPage, bitmap_screen, tolerance)
	// 	if fx != -1 && fy != -1 {
	// 		robotgo.MoveClick(fx, fy, "left", false)
	// 		robotgo.FreeBitmap(bitmap_screen)
	// 		goto RollAngelBeatPool
	// 	}
	// 	robotgo.FreeBitmap(bitmap_screen)
	// 	goto GachaPageCheck

	// E 3 times
	for i := 0; i < 3; i++ {
		robotgo.KeyTap("e")
		robotgo.Sleep(2)
	}

	// gacha-abPool
	findImageOnScreen("3.2", abPool, tolerance, false, false, true)
	// gacha-abPool2
	findImageOnScreen("3.3", abPool2, tolerance, false, false, true)
	// keep enter until rollAgain and click rollAgain
	findImageOnScreen("3.4", rollAgain, tolerance, false, true, true)
	// gacha-abPool2
	findImageOnScreen("3.5", abPool2, tolerance, false, false, true)
	// keep enter until rollAgain
	findImageOnScreen("3.6", rollAgain, tolerance, false, true, false)
	robotgo.KeyTap("esc")
	// wait until gacha-abPool
	findImageOnScreen("3.7", atGachaPage, tolerance, false, false, false)
	// E 3 times
	for i := 0; i < 3; i++ {
		robotgo.KeyTap("e")
		robotgo.Sleep(2)
	}
	// gacha-pinkPool
	findImageOnScreen("3.8", pinkPool, tolerance, false, false, true)
	// gacha-abPool2
	findImageOnScreen("3.9", abPool2, tolerance, false, false, true)
	findImageOnScreen("3.10", gachaResult, tolerance, false, true, true)
	// wait until atGachaPage
	findImageOnScreen("3.11", atGachaPage, tolerance+0.05, false, false, false)
	// E 5 times
	for i := 0; i < 5; i++ {
		robotgo.KeyTap("e")
		robotgo.Sleep(2)
	}
	// gacha-bluePool
	findImageOnScreen("3.12", bluePool, tolerance, false, false, true)
	// gacha-abPool2
	findImageOnScreen("3.13", abPool2, tolerance, false, false, true)
	// keep enter until rollAgain and click rollAgain
	findImageOnScreen("3.14", rollAgain, tolerance, false, true, true)
	// gacha-abPool2
	findImageOnScreen("3.15", abPool2, tolerance, false, false, true)
	findImageOnScreen("3.16", gachaResult, tolerance, false, true, true)
	// wait until atGachaPage
	findImageOnScreen("3.17", atGachaPage, tolerance, false, false, false)
	// esc
	robotgo.KeyTap("esc")

	// wait until mainPage
	// click mainPage (enter?)
	// enter(?)
	// keep esc until mainPage
	findImageOnScreen("3.18", mainPage, tolerance-0.15, true, true, false)
}

func phaseCheck() bool {
	// 進選單
	findImageOnScreen("4.1", mainPage, tolerance, false, false, true)
	// 進小隊頁
	findImageOnScreen("4.2", team, tolerance, false, false, true)
	//進篩選頁
	findImageOnScreen("4.3", teamFilter, tolerance, false, false, true)
	//篩ss & click ok
	findImageOnScreen("4.4", teamFilterSS, tolerance, false, false, true)
	findImageOnScreen("4.5", ok_png, tolerance, false, false, true)

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
	robotgo.KeyTap("esc")
	findImageOnScreen("5.1", settings, tolerance, false, false, true)
	// keep enter until teachHand and click teachHand
	// wait until teachHand
	// wait until keyboardCheck
	// click enter
	// esc
	// 上
	// 上
	// enter
	// OK(or enter)
	findImageOnScreen("5.2", teachHand13, tolerance, false, true, true)
	findImageOnScreen("5.3", teachHand14, tolerance, false, false, true)
	robotgo.Sleep(2)
	robotgo.KeyTap("enter")
	robotgo.Sleep(2)
	robotgo.KeyTap("esc")
	robotgo.Sleep(2)
	robotgo.KeyTap("up")
	findImageOnScreen("5.4", returnFirstPage, tolerance, false, false, true)

	// wait menuOfFirstPage and click it
	findImageOnScreen("5.5", menuOfFirstPage, tolerance, false, false, true)
	// wait deleteAccount and click it
	findImageOnScreen("5.6", deleteAccount, tolerance, false, false, true)
	// wait deleteAccount2 and click it
	findImageOnScreen("5.7", deleteAccount2, tolerance, false, false, true)
	// OK
	findImageOnScreen("5.8", ok_png, tolerance, false, false, true)
	// OK
	findImageOnScreen("5.9", ok_png, tolerance, false, false, true)

	// (接續phaseBegin)

}

func exitEvent() {
	exit_event := robotgo.AddEvent("f1")
	if exit_event {
		fmt.Println("終止程式")
		os.Exit(0)
	}
}

// func errHandler(err error, msg string) {
// 	if err != nil {
// 		fmt.Println(msg)
// 		panic(err)
// 	}
// }
