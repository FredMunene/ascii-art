package functions

import (
	"os"
	"os/exec"
	"testing"
)

type TestCases struct {
	Name     string
	Input    string
	Expected string
}

// func TestCheckBannerName(t *testing.T) {
// 	testCases := []TestCases{
// 		{Name: "provided filename is shadow", Input: "shadow", Expected: "shadow.txt"},
// 		{Name: "provided filename is standard", Input: "standard", Expected: "standard.txt"},
// 		{Name: "provided filename is thinkertoy", Input: "thinkertoy", Expected: "thinkertoy.txt"},
// 		{Name: "provided filename is not shadow, standard,thinkertoy", Input: "othername", Expected: ""},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.Name, func(t *testing.T) {
// 			result := CheckBannerName(tc.Input)
// 			if result != tc.Expected {
// 				t.Errorf("expected %v, got %v", tc.Expected, result)
// 			}
// 		})
// 	}
// }

func TestReadInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		exit     bool
	}{
		{"0123456789abcdefghijklmnopqrstvuvwxyz", "0123456789abcdefghijklmnopqrstvuvwxyz", false},
		{"会意; 會意", "", true},
		{" !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~", " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~", false},
		{"Hello\nWorld", "Hello\\nWorld", false},
	}

	for _, tt := range testCases {
		if tt.exit {
			// Run the test in a separate process if it is expected to exit
			t.Run(tt.input, func(t *testing.T) {
				if os.Getenv("TEST_EXIT") == "1" {
					ReadInput(tt.input)
					return
				}
				// 
				cmd := exec.Command(os.Args[0], "-test.run=TestReadInput")
				cmd.Env = append(os.Environ(), "TEST_EXIT=1")
				err := cmd.Run()
				if e, ok := err.(*exec.ExitError); ok && e.ExitCode() == 1 {
					return
				}
				t.Fatalf("expected process to exit with status 1, but got %v", err)
			})
		} else {
			result := ReadInput(tt.input)
			if result != tt.expected {
				t.Errorf("ReadInput(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		}
	}
}
//  func TestPrintAsciiArt(t *testing.T){

//  }

// func TestReadBannerContent (t *testing.T){
// 	thinkertoyMap :=
// 	standardMap := map[32:[                                                       ] 33:[ _   | |  | |  | |  |_|  (_)           ] 34:[ _ _   ( | )   V V                                     ] 35:[   _  _      _| || |_   |_  __  _|   _| || |_   |_  __  _|    |_||_|                           ] 36:[  _     | |   / __)  \__ \  (   /   |_|                ] 37:[ _   __  (_) / /     / /     / /     / / _   /_/ (_)                   ] 38:[            ___      ( _ )     / _ \/\  | (_>  <   \___/\/                     ] 39:[ _   ( )  |/                           ] 40:[  __   / /  | |   | |   | |   | |    \_\       ] 41:[__    \ \    | |   | |   | |   | |  /_/        ] 42:[    _       /\| |/\    \ ` ' /   |_     _|   / , . \    \/|_|\/                        ] 43:[            _      _| |_   |_   _|    |_|                              ] 44:[                     _   ( )  |/       ] 45:[                     ______   |______|                                         ] 46:[                     _   (_)           ] 47:[     __      / /     / /     / /     / /     /_/                       ] 48:[           ___     / _ \   | | | |  | |_| |   \___/                    ] 49:[      _   / |  | |  | |  |_|           ] 50:[          ____    |___ \     __) |   / __/   |_____|                   ] 51:[          _____   |___ /     |_ \    ___) |  |____/                    ] 52:[           _  _     | || |    | || |_   |__   _|     |_|                       ] 53:[          ____    | ___|   |___ \     __) |  |____/                    ] 54:[           __      / /     | '_ \   | (_) |   \___/                    ] 55:[          _____   |___  |     / /     / /     /_/                      ] 56:[           ___     ( _ )    / _ \   | (_) |   \___/                    ] 57:[           ___     / _ \   | (_) |   \__, |     / /     /_/            ] 58:[      _   (_)        _   (_)           ] 59:[      _   (_)        _   ( )  |/       ] 60:[   __    / /   / /   < <     \ \     \_\               ] 61:[           ______   |______|   ______   |______|                               ] 62:[__     \ \     \ \     > >   / /   /_/                 ] 63:[ ___    |__ \      ) |    / /    |_|     (_)                   ] 64:[              ____      / __ \    / / _` |  | | (_| |   \ \__,_|    \____/             ] 65:[                /\         /  \       / /\ \     / ____ \   /_/    \_\                         ] 66:[ ____    |  _ \   | |_) |  |  _ <   | |_) |  |____/                    ] 67:[  _____    / ____|  | |       | |       | |____    \_____|                     ] 68:[ _____    |  __ \   | |  | |  | |  | |  | |__| |  |_____/                      ] 69:[ ______   |  ____|  | |__     |  __|    | |____   |______|                     ] 70:[ ______   |  ____|  | |__     |  __|    | |       |_|                          ] 71:[  _____    / ____|  | |  __   | | |_ |  | |__| |   \_____|                     ] 72:[ _    _   | |  | |  | |__| |  |  __  |  | |  | |  |_|  |_|                     ] 73:[ _____   |_   _|    | |      | |     _| |_   |_____|                   ] 74:[      _        | |       | |   _   | |  | |__| |   \____/                      ] 75:[ _  __  | |/ /  | ' /   |  <    | . \   |_|\_\                 ] 76:[ _        | |       | |       | |       | |____   |______|                     ] 77:[ __  __   |  \/  |  | \  / |  | |\/| |  | |  | |  |_|  |_|                     ] 78:[ _   _   | \ | |  |  \| |  | . ` |  | |\  |  |_| \_|                   ] 79:[  ____     / __ \   | |  | |  | |  | |  | |__| |   \____/                      ] 80:[ _____    |  __ \   | |__) |  |  ___/   | |       |_|                          ] 81:[  ____     / __ \   | |  | |  | |  | |  | |__| |   \___\_\                     ] 82:[ _____    |  __ \   | |__) |  |  _  /   | | \ \   |_|  \_\                     ] 83:[  _____    / ____|  | (___     \___ \    ____) |  |_____/                      ] 84:[ _______   |__   __|     | |        | |        | |        |_|                          ] 85:[ _    _   | |  | |  | |  | |  | |  | |  | |__| |   \____/                      ] 86:[__      __  \ \    / /   \ \  / /     \ \/ /       \  /         \/                             ] 87:[__          __  \ \        / /   \ \  /\  / /     \ \/  \/ /       \  /\  /         \/  \/                                     ] 88:[__   __  \ \ / /   \ V /     > <     / . \   /_/ \_\                   ] 89:[__     __  \ \   / /   \ \_/ /     \   /       | |        |_|                          ] 90:[ ______  |___  /     / /     / /     / /__   /_____|                   ] 91:[ ___   |  _|  | |    | |    | |    | |_   |___|        ] 92:[__       \ \       \ \       \ \       \ \       \_\                   ] 93:[ ___   |_  |    | |    | |    | |   _| |  |___|        ] 94:[ /\   |/\|                                     ] 95:[                                                             ______   |______| ] 96:[ _   ( )   \|                          ] 97:[                    __ _    / _` |  | (_| |   \__,_|                   ] 98:[ _       | |      | |__    | '_ \   | |_) |  |_.__/                    ] 99:[                  ___    / __|  | (__    \___|                 ] 100:[     _       | |    __| |   / _` |  | (_| |   \__,_|                   ] 101:[                  ___    / _ \  |  __/   \___|                 ] 102:[  __    / _|  | |_   |  _|  | |    |_|                 ] 103:[                    __ _    / _` |  | (_| |   \__, |    __/ |   |___/  ] 104:[ _       | |      | |__    |  _ \   | | | |  |_| |_|                   ] 105:[ _   (_)   _   | |  | |  |_|           ] 106:[   _     (_)     _     | |    | |    | |   _/ |  |__/  ] 107:[         _      | | _   | |/ /  |   <   |_|\_\                 ] 108:[ _   | |  | |  | |  | |  |_|           ] 109:[                           _ __ ___    | '_ ` _ \   | | | | | |  |_| |_| |_|                           ] 110:[                   _ __    | '_ \   | | | |  |_| |_|                   ] 111:[                    ___     / _ \   | (_) |   \___/                    ] 112:[                   _ __    | '_ \   | |_) |  | .__/   | |      |_|     ] 113:[                    __ _    / _` |  | (_| |   \__, |      | |      |_| ] 114:[                 _ __   | '__|  | |     |_|                    ] 115:[               ___   / __|  \__ \  |___/               ] 116:[ _     | |    | |_   | __|  \ |_    \__|               ] 117:[                   _   _   | | | |  | |_| |   \__,_|                   ] 118:[                  __   __  \ \ / /   \ V /     \_/                     ] 119:[                        __      __  \ \ /\ / /   \ V  V /     \_/\_/                           ] 120:[                __  __  \ \/ /   >  <   /_/\_\                 ] 121:[                   _   _   | | | |  | |_| |   \__, |   __/ /   |___/   ] 122:[               ____  |_  /   / /   /___|               ] 123:[   __    / /   | |   / /    \ \     | |     \_\        ] 124:[ _   | |  | |  | |  | |  | |  | |  |_| ] 125:[__     \ \     | |     \ \    / /   | |   /_/          ] 126:[ /\/|  |/\/                                            ]]

// 	shadowMap :=
// 	testCases := []TestCases{
// 		{Name: "banner file is shadow.txt",Input:"shadow.txt",Expected:"0123456789abcdefghijklmnopqrstvuvwxyz"},
// 		{Name: "banner file is standard.txt",Input:"standard.txt",Expected:"standard.txt"},
// 		//{Name: "banner file is thinkertoy.txt",Input:"thinkertoy.txt",Expected:"standard.txt"},
// 		}
// 	for _, tc := range testCases{
// 		t.Run(tc.Name, func(t *testing.T) {
// 			result := ReadInput(tc.Input)
// 			if result != tc.Expected{
// 				t.Fatalf("expected %v, got %v", tc.Expected,result)
// 			}
// 		})
// 	}
// }
