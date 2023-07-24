package main

import (
	"C"
	"fmt"
	"math"

	"github.com/EndlessCheng/mahjong-helper/util"
	"github.com/EndlessCheng/mahjong-helper/util/model"
)

func kanaTileToMHTile(kanaTile int) int {
	switch kanaTile / 10 {
	//m
	case 0:
		//red 5m
		if kanaTile == 0 {
			return 4
		}
		return kanaTile - 1
	//p
	case 1:
		//red 5p
		if kanaTile%10 == 0 {
			return 4 + 9
		}
		return kanaTile - 2
	//s
	case 2:
		//red 5s
		if kanaTile%10 == 0 {
			return 4 + 9*2
		}
		return kanaTile - 3
	//z
	case 3:
		return kanaTile - 3
	}
	fmt.Println("Weird...")
	return 0
}

func kanaDoraTOMHDoraIndicator(kanaDoraIndicators []int) ([]int, int) {
	i := 0
	nums := make([]int, 0)
	for ; kanaDoraIndicators[i] < 203; i++ {
		nums = append(nums, kanaTileToMHTile(kanaDoraIndicators[i]-18-(37*i)))
	}
	return nums, i
}

func kanaHandTOMHTile(kanaHand int) (int, int) {
	offset := 353
	switch (kanaHand - offset) / 36 {
	//m, p, s
	case 0, 1, 2:
		switch (kanaHand - offset) % 36 {
		//red 5
		case 0:
			return (kanaHand - offset) / 36, 4 + ((kanaHand-offset)/36)*9
		//1
		case 1, 2, 3, 4:
			return 4, 0 + ((kanaHand-offset)/36)*9
		//2
		case 5, 6, 7, 8:
			return 4, 1 + ((kanaHand-offset)/36)*9
		//3
		case 9, 10, 11, 12:
			return 4, 2 + ((kanaHand-offset)/36)*9
		//4
		case 13, 14, 15, 16:
			return 4, 3 + ((kanaHand-offset)/36)*9
		//black 5
		case 17, 18, 19:
			return 4, 4 + ((kanaHand-offset)/36)*9
		//6
		case 20, 21, 22, 23:
			return 4, 5 + ((kanaHand-offset)/36)*9
		//7
		case 24, 25, 26, 27:
			return 4, 6 + ((kanaHand-offset)/36)*9
		//8
		case 28, 29, 30, 31:
			return 4, 7 + ((kanaHand-offset)/36)*9
		//9
		case 32, 33, 34, 35:
			return 4, 8 + ((kanaHand-offset)/36)*9
		}
	//z
	case 3:
		return 4, (kanaHand-477)/4 + 27
	}
	fmt.Println("Weird...")
	return 4, 0
}

func kanaHandTOMHCount(kanaHand []int) ([]int, []int) {
	i := 0
	nums := make([]int, 34)
	reds := make([]int, 3)
	for ; i < len(kanaHand); i++ {
		if kanaHand[i] < 489 {
			red, tile := kanaHandTOMHTile(kanaHand[i])
			nums[tile]++
			if red != 4 {
				reds[red] = 1
			}
		} else if kanaHand[i] < 526 {
			tile := kanaTileToMHTile(kanaHand[i] - 489)
			nums[tile]++
		}
	}
	return nums, reds
}

func kanaprogressionTOMHDiscardTile(kanaProgression int) (int, int, bool, bool) {
	offset := 5
	seat := (kanaProgression - offset) / 148
	kanaTile := (kanaProgression - offset - seat*148) / 4
	tile := kanaTileToMHTile(kanaTile)
	moqi := (kanaProgression - offset - seat*148 - kanaTile*4) / 2
	riichi := (kanaProgression - offset - seat*148 - kanaTile*4 - moqi*2)
	return seat, tile, moqi == 1, riichi == 1
}

func kanaProgressionTOMHGlobalDiscardTiles(kanaProgression []int) []int {
	nums := make([]int, 0)
	for _, progression := range kanaProgression {
		if progression < 5 {
			continue
		}
		if progression < 597 {
			_, tile, is_moqi, _ := kanaprogressionTOMHDiscardTile(progression)
			if is_moqi {
				tile = ^tile
			} else {

			}
			nums = append(nums, tile)
		}
	}
	return nums
}

func absGlobalDiscardTiles(globalDiscardTiles []int) []int {
	nums := make([]int, 0)
	copy(nums, globalDiscardTiles)
	normalDiscardTiles(nums)
	// for i, tile := range globalDiscardTiles {
	// 	if tile < 0 {
	// 		nums[i] = ^tile
	// 	}
	// }
	return nums
}

func kanaProgressionTOMeldInfo(kanaProgression []int) ([][]*model.Meld, [][]int, [][]int, []bool, [][]int, []int, [][]int, []bool, []bool, []int, []int) {
	var chiDictNum = [][]int{
		{1, 2, 0},
		{0, 2, 1},
		{2, 3, 1},
		{0, 1, 2},
		{1, 3, 2},
		{3, 4, 2},
		{3, 4, 2},
		{1, 2, 3},
		{2, 4, 3},
		{2, 4, 3},
		{4, 5, 3},
		{4, 5, 3},
		{2, 3, 4},
		{2, 3, 4},
		{3, 5, 4},
		{3, 5, 4},
		{5, 6, 4},
		{5, 6, 4},
		{3, 4, 5},
		{3, 4, 5},
		{4, 6, 5},
		{4, 6, 5},
		{6, 7, 5},
		{4, 5, 6},
		{4, 5, 6},
		{5, 7, 6},
		{7, 8, 6},
		{5, 6, 7},
		{6, 8, 7},
		{6, 7, 8},
	}
	var chiDictNum_sorted = [][]int{
		{0, 1, 2},
		{0, 1, 2},
		{1, 2, 3},
		{0, 1, 2},
		{1, 2, 3},
		{2, 3, 4},
		{2, 3, 4},
		{1, 2, 3},
		{2, 3, 4},
		{2, 3, 4},
		{3, 4, 5},
		{3, 4, 5},
		{2, 3, 4},
		{2, 3, 4},
		{3, 4, 5},
		{3, 4, 5},
		{4, 5, 6},
		{4, 5, 6},
		{3, 4, 5},
		{3, 4, 5},
		{4, 5, 6},
		{4, 5, 6},
		{5, 6, 7},
		{4, 5, 6},
		{4, 5, 6},
		{5, 6, 7},
		{6, 7, 8},
		{5, 6, 7},
		{6, 7, 8},
		{6, 7, 8},
	}
	var chiDictBool = [][]bool{
		{false, false},
		{false, false},
		{false, false},
		{false, false},
		{false, false},
		{false, false},
		{true, false},
		{false, false},
		{false, false},
		{true, false},
		{false, false},
		{true, false},
		{false, false},
		{true, true},
		{false, false},
		{true, true},
		{false, false},
		{true, true},
		{false, false},
		{true, false},
		{false, false},
		{true, false},
		{false, false},
		{false, false},
		{true, false},
		{false, false},
		{false, false},
		{false, false},
		{false, false},
		{false, false},
	}
	var ponDictNum = [][]int{
		{0, 0, 0},
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
		{4, 4, 4},
		{4, 4, 4},
		{4, 4, 4},
		{5, 5, 5},
		{6, 6, 6},
		{7, 7, 7},
		{8, 8, 8},
	}
	var ponDictBool = [][]bool{
		{false, false},
		{false, false},
		{false, false},
		{false, false},
		{false, false},
		{true, false},
		{true, true},
		{false, false},
		{false, false},
		{false, false},
		{false, false},
	}
	melds := make([][]*model.Meld, 4)
	for i := range melds {
		melds[i] = make([]*model.Meld, 0)
	}
	meldsDiscardsAtGlobal := make([][]int, 4)
	for i := range meldsDiscardsAtGlobal {
		meldsDiscardsAtGlobal[i] = make([]int, 0)
	}
	meldsDiscardsAt := make([][]int, 4)
	for i := range meldsDiscardsAt {
		meldsDiscardsAt[i] = make([]int, 0)
	}
	playersDiscardTiles := make([][]int, 4)
	for i := range playersDiscardTiles {
		playersDiscardTiles[i] = make([]int, 0)
	}
	playersLatestDiscardAtGlobal := make([]int, 4)
	for i := range playersLatestDiscardAtGlobal {
		playersLatestDiscardAtGlobal[i] = -1
	}
	playersEarlyOutsideTiles := make([][]int, 4)
	for i := range playersEarlyOutsideTiles {
		playersEarlyOutsideTiles[i] = make([]int, 0)
	}
	playersReachTileAtGlobal := make([]int, 4)
	for i := range playersReachTileAtGlobal {
		playersReachTileAtGlobal[i] = -1
	}
	playersReachTileAt := make([]int, 4)
	for i := range playersReachTileAt {
		playersReachTileAt[i] = -1
	}
	playerIsNaki := make([]bool, 4)
	playerIsReached := make([]bool, 4)
	playerCanIppatsu := make([]bool, 4)
	var seat int
	var mpsz int
	globalDiscardCount := 0
	playerDiscardCount := make([]int, 4)
	for _, progression := range kanaProgression {
		if progression < 5 {
			continue
		}
		//Discard
		if progression < 597 {
			seat, tile, moqi, riichi := kanaprogressionTOMHDiscardTile(progression)
			mtile := 0
			if moqi {
				mtile = ^tile
			} else {
				mtile = tile
			}
			if playerCanIppatsu[seat] {
				playerCanIppatsu[seat] = false
			}
			if riichi {
				playerIsReached[seat] = true
				playerCanIppatsu[seat] = true
				playersReachTileAtGlobal[seat] = globalDiscardCount
				playersReachTileAt[seat] = playerDiscardCount[seat]
			}
			playersDiscardTiles[seat] = append(playersDiscardTiles[seat], mtile)
			if playerDiscardCount[seat] < 5 && !playerIsReached[seat] {
				playersEarlyOutsideTiles[seat] = append(playersEarlyOutsideTiles[seat], util.OutsideTiles(tile)...)
			}
			playersLatestDiscardAtGlobal[seat] = globalDiscardCount
			playerDiscardCount[seat]++
			globalDiscardCount++

			continue
		}
		if progression < 2165 {
			playerCanIppatsu = []bool{false, false, false, false}
			switch {
			//Chi
			case 597 <= progression && progression < 957:
				var myMeld model.Meld
				seat = (progression - 597) / 90
				myMeld.MeldType = meldTypeChi
				chiType := ((progression - 597) % 90) % 30
				mpsz = ((progression - 597) % 90) / 30

				myMeld.CalledTile = chiDictNum[chiType][2] + mpsz*9
				myMeld.SelfTiles = []int{
					chiDictNum[chiType][0] + mpsz*9,
					chiDictNum[chiType][1] + mpsz*9,
				}
				myMeld.Tiles = []int{
					chiDictNum_sorted[chiType][0] + mpsz*9,
					chiDictNum_sorted[chiType][1] + mpsz*9,
					chiDictNum_sorted[chiType][2] + mpsz*9,
				}
				myMeld.ContainRedFive = chiDictBool[chiType][0]
				myMeld.RedFiveFromOthers = chiDictBool[chiType][1]

				melds[seat] = append(melds[seat], &myMeld)
				meldsDiscardsAtGlobal[seat] = append(meldsDiscardsAtGlobal[seat], globalDiscardCount)
				meldsDiscardsAt[seat] = append(meldsDiscardsAt[seat], playerDiscardCount[seat])
				playerIsNaki[seat] = true
			//Pon
			case 957 <= progression && progression < 1437:
				var myMeld model.Meld
				seat = (progression - 957) / 120
				myMeld.MeldType = meldTypePon
				ponType := ((progression - 957) % 40) % 11
				mpsz = ((progression - 957) % 40) / 11
				if mpsz < 3 {
					myMeld.CalledTile = ponDictNum[ponType][2] + mpsz*9
					myMeld.SelfTiles = []int{
						ponDictNum[ponType][0] + mpsz*9,
						ponDictNum[ponType][1] + mpsz*9,
					}
					myMeld.Tiles = []int{
						ponDictNum[ponType][0] + mpsz*9,
						ponDictNum[ponType][1] + mpsz*9,
						ponDictNum[ponType][2] + mpsz*9,
					}
					myMeld.ContainRedFive = ponDictBool[ponType][0]
					myMeld.RedFiveFromOthers = ponDictBool[ponType][1]
				} else {
					myMeld.CalledTile = ponType + mpsz*9
					myMeld.SelfTiles = []int{
						ponType + mpsz*9,
						ponType + mpsz*9,
					}
					myMeld.Tiles = []int{
						ponType + mpsz*9,
						ponType + mpsz*9,
						ponType + mpsz*9,
					}
					myMeld.ContainRedFive = false
					myMeld.RedFiveFromOthers = false
				}
				melds[seat] = append(melds[seat], &myMeld)
				meldsDiscardsAtGlobal[seat] = append(meldsDiscardsAtGlobal[seat], globalDiscardCount)
				meldsDiscardsAt[seat] = append(meldsDiscardsAt[seat], playerDiscardCount[seat])
				playerIsNaki[seat] = true
			//Ming Gang
			case 1437 <= progression && progression < 1881:
				var myMeld model.Meld
				seat = (progression - 1437) / 111
				myMeld.MeldType = meldTypeMinkan
				minkanType := ((progression - 1437) % 111) % 37
				mpsz = (((progression - 1437) % 111) % 37) / 10

				myMeld.CalledTile = kanaTileToMHTile(minkanType)
				myMeld.SelfTiles = []int{
					myMeld.CalledTile,
					myMeld.CalledTile,
					myMeld.CalledTile,
				}
				myMeld.Tiles = []int{
					myMeld.CalledTile,
					myMeld.CalledTile,
					myMeld.CalledTile,
					myMeld.CalledTile,
				}
				myMeld.ContainRedFive = false
				myMeld.RedFiveFromOthers = false
				if mpsz < 3 {
					if minkanType%10 == 0 {
						myMeld.ContainRedFive = true
						myMeld.RedFiveFromOthers = true
					}
					if minkanType%10 == 5 {
						myMeld.ContainRedFive = true
						myMeld.RedFiveFromOthers = false
					}
				}
				melds[seat] = append(melds[seat], &myMeld)
				meldsDiscardsAtGlobal[seat] = append(meldsDiscardsAtGlobal[seat], globalDiscardCount)
				meldsDiscardsAt[seat] = append(meldsDiscardsAt[seat], playerDiscardCount[seat])
				playerIsNaki[seat] = true
			//An Gang
			case 1881 <= progression && progression < 2017:
				var myMeld model.Meld
				seat = (progression - 1881) / 34
				myMeld.MeldType = meldTypeAnkan
				ankanType := ((progression - 1881) % 34)
				mpsz = ((progression - 1881) % 34) / 9

				myMeld.CalledTile = ankanType
				myMeld.SelfTiles = []int{
					myMeld.CalledTile,
					myMeld.CalledTile,
					myMeld.CalledTile,
				}
				myMeld.Tiles = []int{
					myMeld.CalledTile,
					myMeld.CalledTile,
					myMeld.CalledTile,
					myMeld.CalledTile,
				}
				myMeld.ContainRedFive = false
				myMeld.RedFiveFromOthers = false
				if mpsz < 3 {
					if ankanType%9 == 4 {
						myMeld.ContainRedFive = true
						myMeld.RedFiveFromOthers = false
					}
				}
				melds[seat] = append(melds[seat], &myMeld)
				meldsDiscardsAtGlobal[seat] = append(meldsDiscardsAtGlobal[seat], globalDiscardCount)
				meldsDiscardsAt[seat] = append(meldsDiscardsAt[seat], playerDiscardCount[seat])
			//Jia Gang
			case 2017 <= progression && progression < 2165:
				var myMeld model.Meld
				seat = (progression - 2017) / 37
				myMeld.MeldType = meldTypeKakan
				kakanType := ((progression - 2017) % 37)
				mpsz = ((progression - 2017) % 37) / 10

				myMeld.CalledTile = kanaTileToMHTile(kakanType)
				myMeld.SelfTiles = []int{
					myMeld.CalledTile,
					myMeld.CalledTile,
					myMeld.CalledTile,
				}
				myMeld.Tiles = []int{
					myMeld.CalledTile,
					myMeld.CalledTile,
					myMeld.CalledTile,
					myMeld.CalledTile,
				}
				myMeld.ContainRedFive = false
				myMeld.RedFiveFromOthers = false
				if mpsz < 3 {
					if kakanType%9 == 4 {
						myMeld.ContainRedFive = true
						myMeld.RedFiveFromOthers = false
					}
				}
				for _, _meld := range melds[seat] {
					if _meld.Tiles[0] == myMeld.CalledTile && _meld.MeldType == meldTypePon {
						_meld.MeldType = meldTypeKakan
						_meld.Tiles = myMeld.Tiles
						_meld.ContainRedFive = myMeld.ContainRedFive
						break
					}
				}
				playerIsNaki[seat] = true
				fmt.Println("kakan not fount?")
			}
		}
	}
	return melds, meldsDiscardsAtGlobal, meldsDiscardsAt, playerIsNaki, playersDiscardTiles, playersLatestDiscardAtGlobal, playersEarlyOutsideTiles, playerIsReached, playerCanIppatsu, playersReachTileAtGlobal, playersReachTileAt
}

const (
	action_DiscardTile = iota
	action_Riichi
	action_AnKan
	action_Kakan
	action_Zimo
	action_Liuju
	action_Cancel
	action_Chi
	action_Pon
	action_MinKan
	action_Ron
)

func actualAction(options []int, action_index int) (int, int) {
	action := options[action_index]
	action_type := -1
	targetTile := -1

	var chiDictNum = [][]int{
		{1, 2, 0},
		{0, 2, 1},
		{2, 3, 1},
		{0, 1, 2},
		{1, 3, 2},
		{3, 4, 2},
		{3, 4, 2},
		{1, 2, 3},
		{2, 4, 3},
		{2, 4, 3},
		{4, 5, 3},
		{4, 5, 3},
		{2, 3, 4},
		{2, 3, 4},
		{3, 5, 4},
		{3, 5, 4},
		{5, 6, 4},
		{5, 6, 4},
		{3, 4, 5},
		{3, 4, 5},
		{4, 6, 5},
		{4, 6, 5},
		{6, 7, 5},
		{4, 5, 6},
		{4, 5, 6},
		{5, 7, 6},
		{7, 8, 6},
		{5, 6, 7},
		{6, 8, 7},
		{6, 7, 8},
	}
	var ponDictNum = []int{
		0, 1, 2, 3, 4, 4, 4, 5, 6, 7, 8,
		9, 10, 11, 12, 13, 13, 13, 14, 15, 16, 17,
		18, 19, 20, 21, 22, 22, 22, 23, 24, 25, 26,
		27, 28, 29, 30, 31, 32, 33,
	}

	if 0 <= action && action <= 147 {
		action_type = action_DiscardTile
		targetTile = kanaTileToMHTile((action) / 4)
		if action%2 == 1 {
			action_type = action_Riichi
		}
	}
	if 148 <= action && action <= 181 {
		action_type = action_AnKan
		targetTile = action - 148
	}
	if 182 <= action && action <= 218 {
		action_type = action_Kakan
		targetTile = kanaTileToMHTile(action - 148)
	}
	if 219 == action {
		action_type = action_Zimo
	}
	if 220 == action {
		action_type = action_Liuju
	}
	if 221 == action {
		action_type = action_Cancel
	}
	if 222 <= action && action <= 311 {
		action_type = action_Chi
		targetTile = chiDictNum[action-222][2]
	}
	// relSeatName = ['下家', '對家', '上家']
	if 312 <= action && action <= 431 {
		action_type = action_Pon
		targetTile = ponDictNum[(action-312)%40]
	}
	if 432 <= action && action <= 542 {
		action_type = action_MinKan
		targetTile = kanaTileToMHTile((action - 432) % 37)
	}
	if 543 <= action && action <= 545 {
		action_type = action_Ron
	}
	if 546 <= action || action <= -1 {
	}
	return action_type, targetTile
}

func Max(array []int) int {
	var max int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

//export getReward
func getReward(sparse []int, numeric []int, progression []int, option []int, action []int) int {
	// //FIRST SPARSE FEATURES
	// var sparse []int
	// sparse = []int{4, 6, 9, 11, 15, 24, 208, 286, 291, 308, 309, 328, 332, 345, 350, 353, 362, 366, 377, 381, 389, 394, 395, 409, 413, 442, 445, 449, 515}
	// //FIRST NUMERIC FEATURES
	// var numeric []int
	// numeric = []int{0, 0, 24000, 29000, 23000, 24000}
	// //FIRST PROGRESSION FEATURES
	// var progression []int
	// progression = []int{0, 157, 345, 581, 131, 281, 425, 569, 131, 297, 421, 593, 149, 277, 373, 583, 145, 265, 375, 523, 53, 173, 345, 565, 61, 241, 347, 519, 41, 253, 409, 952, 493, 19, 217, 391, 515, 125, 167, 335, 499, 15, 1318, 465, 145, 209, 319, 455, 135, 281, 417, 511, 147, 269, 423, 519, 113, 773, 293, 447, 535, 101, 237, 357, 587, 141, 207, 435, 527, 91, 289}
	// //FIRST OPTION FEATURES
	// var option []int
	// option = []int{0, 12, 16, 28, 32, 40, 48, 64, 68, 100, 104, 105, 106, 107, 108}
	// //ACTION INDEX
	// var action []int
	// action = []int{12}
	//17,1000,1000,-3000,1000,2,0,3,1,0,41900,152

	// var p []*playerInfo
	// var d roundData
	// d.gameMode = gameModeMatch
	// d.playerNumber = 4
	// // 场数（如东1为0，东2为1，...，南1为4，...）
	// d.roundNumber = 0
	// // 本场数，从 0 开始算
	// d.benNumber = 0
	// // 场风
	// d.roundWindTile = 27 + d.roundNumber/d.playerNumber
	// // 庄家 0=自家, 1=下家, 2=对家, 3=上家
	// // 请用 reset 设置
	// d.dealer = 0
	// // 宝牌指示牌
	// d.doraIndicators = []int{0}
	// // 自家手牌
	// d.counts = []int{0}
	// // 按照 mps 的顺序记录自家赤5数量，包含副露的赤5
	// // 比如有 0p 和 0s 就是 [1, 0, 1]
	// d.numRedFives = []int{0, 0, 0}
	// // 牌山剩余牌量
	// d.leftCounts = util.InitLeftTiles34WithTiles34(util.MustStrToTiles34("2222333377779999m 22228888p 333355557777s 4444z"))
	// // 全局舍牌
	// // 按舍牌顺序，负数表示摸切(-)，非负数表示手切(+)
	// // 可以理解成：- 表示不要/暗色，+ 表示进张/亮色
	// d.globalDiscardTiles = []int{0}
	// // 0=自家, 1=下家, 2=对家, 3=上家
	// d.players = players: []*playerInfo{
	// 		newPlayerInfo("自家", playerWindTile[0]),
	// 		newPlayerInfo("下家", playerWindTile[1]),
	// 		newPlayerInfo("对家", playerWindTile[2]),
	// 		newPlayerInfo("上家", playerWindTile[3]),
	// 	},

	// fmt.Println("Helloworld")
	debugMode = true

	// d := newGame(&majsoulRoundData{})
	// handsTiles34, _, err := util.StrToTiles34("123456789m 123456789p 123456789s 1234567z")
	// if err != nil {
	// 	fmt.Print("Error: ")
	// 	fmt.Println(err)
	// 	// t.Fatal(err)
	// }
	// globalDiscardTiles34, _, err := util.StrToTiles34("22m 158p 123789s 6z")
	// if err != nil {
	// 	fmt.Print("Error: ")
	// 	fmt.Println(err)
	// 	// t.Fatal(err)
	// }
	// for i, c := range handsTiles34 {
	// 	if c == 0 {
	// 		continue
	// 	}
	// 	d.leftCounts[i] -= c
	// 	if d.leftCounts[c] < 0 {
	// 		fmt.Print("参数有误: ")
	// 		fmt.Println(util.Mahjong[c])
	// 		// t.Fatal("参数有误: ", util.Mahjong[c])
	// 	}
	// }
	// for i, c := range globalDiscardTiles34 {
	// 	if c == 0 {
	// 		continue
	// 	}
	// 	d.leftCounts[i] -= c
	// 	if d.leftCounts[c] < 0 {
	// 		fmt.Print("参数有误: ")
	// 		fmt.Println(util.Mahjong[c])
	// 		// t.Fatal("参数有误: ", util.Mahjong[c])
	// 	}
	// 	d.globalDiscardTiles = append(d.globalDiscardTiles, i)
	// }

	kanaSeatToMH := []int{0, 3, 2, 1}

	d := newGame(&majsoulRoundData{})
	d.gameMode = gameModeMatch
	d.skipOutput = false
	d.playerNumber = 4
	d.dealer = 0
	d.roundWindTile = sparse[3] + 27 - 11
	d.roundNumber = sparse[4] - 14 + (sparse[3]-11)*4
	d.benNumber = numeric[0]
	d.dealer = kanaSeatToMH[sparse[2]-7]
	doraCount := 1
	d.doraIndicators, doraCount = kanaDoraTOMHDoraIndicator(sparse[5:10])
	d.counts, d.numRedFives = kanaHandTOMHCount(sparse[5+doraCount+9:])
	d.globalDiscardTiles = kanaProgressionTOMHGlobalDiscardTiles(progression)
	d.leftCounts = model.InitLeftTiles34WithTiles34(util.TilesToTiles34(append(absGlobalDiscardTiles(d.globalDiscardTiles), util.Tiles34ToTiles(d.counts)...)))

	//From prog: 0123
	//Seat: 0
	//to Player: 0123
	//Seat: 1
	//to Player: 3012
	//Seat: 2
	//to Player: 2301
	//Seat: 3
	//to Player: 1230
	seatSubToObj := [][]int{
		{0, 1, 2, 3},
		{3, 0, 1, 2},
		{2, 3, 0, 1},
		{1, 2, 3, 0},
	}
	playersMelds, playersMeldDiscardsAtGlobal, playersMeldDiscardsAt, playersIsNaki, playersDiscardTiles, playersLatestDiscardAtGlobal, playersEarlyOutsideTiles, playerIsReached, playerCanIppatsu, playersReachTileAtGlobal, playersReachTileAt := kanaProgressionTOMeldInfo(progression)
	d.players[0].selfWindTile = sparse[2] + 27 - 7
	for i := 0; i < 4; i++ {
		d.players[i].selfWindTile = (sparse[2]+i)%4 + 27 - 7
		d.players[seatSubToObj[sparse[2]-7][i]].melds = playersMelds[i]
		d.players[seatSubToObj[sparse[2]-7][i]].meldDiscardsAtGlobal = playersMeldDiscardsAtGlobal[i]
		d.players[seatSubToObj[sparse[2]-7][i]].meldDiscardsAt = playersMeldDiscardsAt[i]
		d.players[seatSubToObj[sparse[2]-7][i]].isNaki = playersIsNaki[i]
		d.players[seatSubToObj[sparse[2]-7][i]].discardTiles = playersDiscardTiles[i]
		d.players[seatSubToObj[sparse[2]-7][i]].latestDiscardAtGlobal = playersLatestDiscardAtGlobal[i]
		d.players[seatSubToObj[sparse[2]-7][i]].earlyOutsideTiles = playersEarlyOutsideTiles[i]
		d.players[seatSubToObj[sparse[2]-7][i]].isReached = playerIsReached[i]
		d.players[seatSubToObj[sparse[2]-7][i]].canIppatsu = playerCanIppatsu[i]
		d.players[seatSubToObj[sparse[2]-7][i]].reachTileAtGlobal = playersReachTileAtGlobal[i]
		d.players[seatSubToObj[sparse[2]-7][i]].reachTileAt = playersReachTileAt[i]
	}
	kanaPlayerInfo := d.newModelPlayerInfo()

	action_type, targetTile := actualAction(option, action[0])

	reward := 0
	rank := 4.0

	for _, v := range numeric[3:] {
		if numeric[2] > v {
			rank -= 1
		}
		if numeric[2] == v {
			rank -= 0.5
		}
	}

	reward += int((2.5 - rank) * 2)

	switch action_type {
	case action_DiscardTile:
		// fmt.Println(targetTile)
		// fmt.Println()
		//Attack
		shanten, results14, incShantenResults14 := util.CalculateShantenWithImproves14(kanaPlayerInfo)
		// fmt.Println(shanten)
		// fmt.Println()

		var action_result util.Hand14AnalysisResult
		decision_index := 0

		for i, r := range incShantenResults14 {
			if r.DiscardTile == targetTile {
				action_result = *r
				decision_index = len(results14) + i
			}
		}
		for i, r := range results14 {
			if r.DiscardTile == targetTile {
				action_result = *r
				decision_index = i
			}
		}

		//Defence
		riskTables := d.analysisTilesRisk()
		mixedRiskTable := riskTables.mixedRiskTable()
		// fmt.Println(mixedRiskTable)

		//聽牌：
		if shanten == 0 {
			reward += 1
		}
		//和牌：
		if shanten == -1 {
			reward += 2
		}
		//鎮聽
		if action_result.Result13.Shanten >= 0 && action_result.Result13.Shanten <= 1 {
			if action_result.Result13.FuritenRate > 0 {
				if action_result.Result13.FuritenRate < 1 {
					// s += "[可能振听]"
					reward -= 1
				} else {
					// s += "[振听]"
					reward -= 2
				}
			}
		}
		//收支能超越第一
		if numeric[2]+int(math.Round(action_result.Result13.MixedRoundPoint)) > Max(numeric[2:]) {
			reward += 1
		}
		//做出的選擇在助手給的選擇排序中佔前四
		if decision_index < 4 {
			reward += 4
		}
		if decision_index > 10 {
			reward -= 2
		}

		defenceRank := 0

		for i := 0; i < 34; i++ {
			if kanaPlayerInfo.HandTiles34[i] != 0 {
				if mixedRiskTable[i] < mixedRiskTable[targetTile] {
					defenceRank += 1
				}
			}
		}

		//防守
		if defenceRank < 3 {
			reward += 1
		}
		if defenceRank < 6 {
			reward += 1
		}
		if defenceRank > 7 {
			reward -= 1
		}
		if defenceRank > 10 {
			reward -= 1
		}

	case action_AnKan:
		reward += 1

	case action_Cancel:
		reward += 1

	case action_Chi:
		reward += 1

	case action_Kakan:
		reward += 1

	case action_Liuju:
		reward += 1

	case action_MinKan:
		reward += 1

	case action_Pon:
		reward += 1

	case action_Riichi:
		reward += 3

	case action_Ron:
		reward += 10

	case action_Zimo:
		reward += 12

	}

	// bestAttackDiscardTile := simpleBestDiscardTile(kanaPlayerInfo)
	// bestDefenceDiscardTile := mixedRiskTable.getBestDefenceTile(playerInfo.HandTiles34)
	return reward
}

func main() {

}
