package consts

type SoundType uint8

const (
	SoundGameIntro SoundType = iota
	SoundGameEnd
	SoundPlayerTankMove
	SoundPlayerTankStay
	SoundPlayerTankCrash
	SoundEnemyTankCrash
	SoundBrickWallCrash
	SoundTankShoot
)
