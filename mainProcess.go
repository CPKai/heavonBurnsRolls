package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/go-vgo/robotgo"
)

var (
	screenWidth     int     = 1920
	screenLength    int     = 1080
	tolerance       float64 = 0.2
	chooseLang      string  = "img/step-chooseLang.png"
	nextPng         string  = "img/step-next.png"
	endTuPng        string  = "img/step-endTutorial.png"
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
	roleSS          string  = "img/role-ss.png"
)

func main() {
	fmt.Println("Heaven Burns Red 工具箱")
	fmt.Println("按f1停止程式")
	fmt.Println("開始進行首抽刷取")
	go exitEvent()
	rollGacha()
	// calculateScore(loadConfig("roleConfig.txt"))
}

func rollGacha() {
	var (
		round_counter int = 0
	)

	for {
		round_counter++
		fmt.Printf("current round:%d\n", round_counter)
		phaseBegin()
		phaseStory()
		phaseGacha()
		result := phaseCheck()
		if result > 1099 {
			fmt.Printf("總分:%v\n", result)
			fmt.Println("總分達到標準，停止程式")
			break
		} else {
			fmt.Printf("總分:%v\n", result)
			fmt.Println("總分未及標準，進行下一輪")
		}
		phaseDel()
	}

}

func loadConfig(csvPath string) map[string]int {

	dataMap := make(map[string]int, 13)
	csvFile, err := os.Open(csvPath)
	errHandler(err, "ErrCode 003001. Open CSV file failed.")

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	errHandler(err, "ErrCode 003002. Read CSV file failed.")

	for _, line := range csvLines {
		tempInt, err := strconv.ParseInt(line[1], 10, 64)
		errHandler(err, "ErrCode 003003. Convert string to int error.")
		dataMap[line[0]] = int(tempInt)
	}

	return dataMap
}

func errHandler(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}

func findImageOnScreen(phaseNum string, imgPath string, sp_tolerance float64, useESC bool, useEnter bool, clickTarget bool) {
	fmt.Printf("phaseNum:%v\t", phaseNum)
	fmt.Printf("imgPath:%s\n", imgPath)
	for {
		if useESC {
			robotgo.KeyTap("esc")
			robotgo.Sleep(1)
		}
		if useEnter {
			robotgo.KeyTap("enter")
			robotgo.Sleep(1)
		}
		robotgo.KeyDown("lctrl")
		robotgo.Sleep(1)
		robotgo.KeyUp("lctrl")
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
func findImageOnScreen_gachaVer(phaseNum string, imgPath string, sp_tolerance float64, useESC bool, useEnter bool, clickTarget bool) {
	fmt.Printf("phaseNum:%v\t", phaseNum)
	fmt.Printf("imgPath:%s\n", imgPath)
	for {
		if useESC {
			robotgo.KeyTap("esc")
			// robotgo.Sleep(1)
		}
		if useEnter {
			robotgo.KeyTap("enter")
			// robotgo.Sleep(1)
		}
		robotgo.KeyDown("lctrl")
		// robotgo.Sleep(1)
		robotgo.KeyUp("lctrl")
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
	// wait 繁體中文 and click
	findImageOnScreen("1.1", chooseLang, tolerance, false, false, true)
}

func phaseStory() {
	findImageOnScreen("2.1", nextPng, tolerance, false, true, true)
	// robotgo.Sleep(1)
	findImageOnScreen("2.2", endTuPng, tolerance, false, true, false)
	robotgo.KeyTap("esc")

	// step - firstGacha
	findImageOnScreen("2.3", firstGacha, tolerance, false, true, true)

	// stop - firstGacha2
	robotgo.Sleep(2)
	findImageOnScreen("2.4", firstGacha2, tolerance, false, false, true)
	robotgo.Sleep(2)

	findImageOnScreen("2.5", teachHand1, tolerance, true, true, true)
	findImageOnScreen("2.6", teachHand2, tolerance, false, false, true)
	findImageOnScreen("2.7", teachHand3, tolerance, false, false, true)
	findImageOnScreen("2.8", teachHand4, tolerance, false, false, true)
	robotgo.Sleep(2)
	robotgo.KeyTap("enter")
	findImageOnScreen("2.9", teachHand5, tolerance, false, false, true)
	findImageOnScreen("2.10", teachHand6, tolerance, false, false, true)
	// wait storyChoose
	findImageOnScreen("2.11", storyChoose, tolerance, true, true, true)
	findImageOnScreen("2.12", teachHand7, tolerance, false, true, true)
	findImageOnScreen("2.13", teachHand8, tolerance, false, false, true)
	findImageOnScreen("2.14", teachHand9, tolerance, false, false, true)
	robotgo.Sleep(2)
	findImageOnScreen("2.15", teachHand10, tolerance, false, false, true)
	robotgo.Sleep(2)
	findImageOnScreen("2.16", teachHand11, tolerance, false, false, true)
	findImageOnScreen("2.17", teachHand12, tolerance, false, false, true)

}

func phaseGacha() {
	findImageOnScreen("3.1", atGachaPage, tolerance, false, false, false)
	// E 3 times
	for i := 0; i < 3; i++ {
		robotgo.KeyTap("e")
		robotgo.Sleep(2)
	}

	// gacha-abPool
	findImageOnScreen_gachaVer("3.2", abPool, tolerance, false, false, true)
	// gacha-abPool2
	findImageOnScreen_gachaVer("3.3", abPool2, tolerance, false, false, true)
	// keep enter until rollAgain and click rollAgain
	findImageOnScreen_gachaVer("3.4", rollAgain, tolerance, false, true, true)
	// gacha-abPool2
	findImageOnScreen_gachaVer("3.5", abPool2, tolerance, false, false, true)
	// keep enter until rollAgain
	findImageOnScreen_gachaVer("3.6", rollAgain, tolerance, false, true, false)
	robotgo.KeyTap("esc")
	// wait until gacha-abPool
	findImageOnScreen_gachaVer("3.7", atGachaPage, tolerance, false, false, false)
	// E 3 times
	for i := 0; i < 3; i++ {
		robotgo.KeyTap("e")
		robotgo.Sleep(2)
	}
	// gacha-pinkPool
	findImageOnScreen_gachaVer("3.8", pinkPool, tolerance, false, false, true)
	// gacha-abPool2
	findImageOnScreen_gachaVer("3.9", abPool2, tolerance, false, false, true)
	findImageOnScreen_gachaVer("3.10", gachaResult, tolerance, false, true, true)
	robotgo.Sleep(1)
	findImageOnScreen_gachaVer("3.11", gachaResult, tolerance, false, true, true)
	// wait until atGachaPage
	findImageOnScreen_gachaVer("3.12", atGachaPage, tolerance, false, false, false)
	// E 5 times
	for i := 0; i < 5; i++ {
		robotgo.KeyTap("e")
		robotgo.Sleep(2)
	}
	// gacha-bluePool
	findImageOnScreen_gachaVer("3.13", bluePool, tolerance, false, false, true)
	// gacha-abPool2
	findImageOnScreen_gachaVer("3.14", abPool2, tolerance, false, false, true)
	// keep enter until rollAgain and click rollAgain
	findImageOnScreen_gachaVer("3.15", rollAgain, tolerance, false, true, true)
	// gacha-abPool2
	findImageOnScreen_gachaVer("3.16", abPool2, tolerance, false, false, true)
	findImageOnScreen_gachaVer("3.17", gachaResult, tolerance, false, true, true)
	robotgo.Sleep(1)
	findImageOnScreen_gachaVer("3.18", gachaResult, tolerance, false, true, true)
	// wait until atGachaPage
	findImageOnScreen_gachaVer("3.19", atGachaPage, tolerance, false, false, false)
	// esc
	robotgo.KeyTap("esc")

	// wait until mainPage
	// click mainPage (enter?)
	// enter(?)
	// keep esc until mainPage
	findImageOnScreen_gachaVer("3.20", mainPage, tolerance, false, false, false)
	fmt.Printf("PhaseNum:3.21\timgPath:%s\n", mainPage)
	for {
		robotgo.KeyTap("enter")
		bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)
		fx, fy := robotgo.FindPic(mainPage, bitmap_screen, tolerance)
		if fx == -1 && fy == -1 {
			robotgo.FreeBitmap(bitmap_screen)
			break
		}
		robotgo.FreeBitmap(bitmap_screen)
	}
	robotgo.Sleep(2)
}

func phaseCheck() int {
	// 進選單
	findImageOnScreen_gachaVer("4.1", mainPage, tolerance, true, true, true)
	// 進小隊頁
	findImageOnScreen("4.2", team, tolerance, false, false, true)
	//進篩選頁
	findImageOnScreen("4.3", teamFilter, tolerance, false, false, true)
	//篩ss & click ok
	findImageOnScreen("4.4", teamFilterSS, tolerance, false, false, true)
	robotgo.Sleep(1)
	findImageOnScreen("4.5", ok_png, tolerance, false, false, true)
	findImageOnScreen("4.6", roleSS, tolerance, false, false, false)

	fmt.Println("phaseNum:4.7\tcalculateScore")
	var score int = calculateScore(loadConfig("roleConfig.txt"))

	return score
}

func calculateScore(roleMap map[string]int) int {
	var score int = 0

	//取得screenshot
	bitmap_screen := robotgo.CaptureScreen(0, 0, screenWidth, screenLength)

	// 把roleConfig中每一組role做對照，確認有無並計分
	for roleName, roleScore := range roleMap {
		fmt.Printf("roleName:%v\t", roleName)
		fmt.Printf("roleScore:%v\t", roleScore)

		if roleName == "gacha-SS_pos_5" {
			fx, fy := robotgo.FindPic("img/"+roleName+".png", bitmap_screen, tolerance)
			if fx == -1 && fy == -1 {
				fmt.Printf("at least 5 ss\n")
				score = score + roleScore
			} else {
				fmt.Printf("lower then 5 ss\n")
			}
		} else {
			fx, fy := robotgo.FindPic("img/"+roleName+".png", bitmap_screen, tolerance+0.1)
			if fx != -1 && fy != -1 {
				fmt.Printf("true\n")
				score = score + roleScore
			} else {
				fmt.Printf("false\n")
			}
		}
	}

	// release memory
	robotgo.FreeBitmap(bitmap_screen)

	return score
}

func phaseDel() {
	robotgo.KeyTap("esc")
	findImageOnScreen_gachaVer("5.1", settings, tolerance, false, false, true)
	// keep enter until teachHand and click teachHand
	// wait until teachHand
	// wait until keyboardCheck
	// click enter
	// esc
	// 上
	// 上
	// enter
	// OK(or enter)
	findImageOnScreen_gachaVer("5.2", teachHand13, tolerance, false, true, true)
	findImageOnScreen_gachaVer("5.3", teachHand14, tolerance, false, false, true)
	robotgo.Sleep(2)
	robotgo.KeyTap("enter")
	robotgo.Sleep(2)
	robotgo.KeyTap("esc")
	robotgo.Sleep(2)
	robotgo.KeyTap("up")
	findImageOnScreen_gachaVer("5.4", returnFirstPage, tolerance, false, false, true)
	findImageOnScreen_gachaVer("5.5", ok_png, tolerance, false, false, true)
	// wait menuOfFirstPage and click it
	findImageOnScreen_gachaVer("5.6", menuOfFirstPage, tolerance, false, false, true)
	// wait deleteAccount and click it
	findImageOnScreen_gachaVer("5.7", deleteAccount, tolerance, false, false, true)
	// wait deleteAccount2 and click it
	findImageOnScreen_gachaVer("5.8", deleteAccount2, tolerance, false, false, true)
	// OK
	findImageOnScreen_gachaVer("5.9", ok_png, tolerance, false, false, true)
	// OK
	findImageOnScreen_gachaVer("5.10", ok_png, tolerance, false, false, true)

	// (接續phaseBegin)

}

func exitEvent() {
	exit_event := robotgo.AddEvent("f1")
	if exit_event {
		fmt.Println("終止程式")
		os.Exit(0)
	}
}
