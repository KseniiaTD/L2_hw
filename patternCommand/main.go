package main

func main() {
	tv := Tv{}
	sound := Sound{}

	onTV := OnCommand{device: &tv}
	onSound := OnCommand{device: &sound}

	offTV := OffCommand{device: &tv}
	offSound := OffCommand{device: &sound}

	onBtnTV := Button{
		command: &onTV,
	}

	offBtnTV := Button{
		command: &offTV,
	}

	onBtnSound := Button{
		command: &onSound,
	}

	offBtnSound := Button{
		command: &offSound,
	}

	onBtnTV.press()
	onBtnSound.press()
	offBtnSound.press()
	offBtnTV.press()

}
