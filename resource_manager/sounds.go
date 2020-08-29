package resource_manager

import (
	"os"
	"path"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/pkg/errors"
	"github.com/shamanr/battle_citty/consts"
)

const (
	soundPath            = "resources/sounds"
	soundGameIntro       = "1 - Track 1.mp3"
	soundGameEnd         = "SFX 1.mp3"
	soundPlayerTankMove  = "3 - Track 3.mp3"
	soundPlayerTankStay  = "4 - Track 4.mp3"
	soundPlayerTankCrash = "SFX 5.mp3"
	soundEnemyTankCrash  = "SFX 6.mp3"
	soundBrickWallCrash  = "SFX 8.mp3"
	soundTankShoot       = "SFX 13.mp3"
)

type SoundInterface interface {
	Init()
	Play()
}

type sound struct {
	soundType consts.SoundType
	soundFile string
	streamer  *beep.StreamSeekCloser
	format    *beep.Format
}

var soundsMap = map[consts.SoundType]*sound{
	consts.SoundGameIntro:       newSound(soundGameIntro),
	consts.SoundGameEnd:         newSound(soundGameEnd),
	consts.SoundPlayerTankMove:  newSound(soundPlayerTankMove),
	consts.SoundPlayerTankStay:  newSound(soundPlayerTankStay),
	consts.SoundPlayerTankCrash: newSound(soundPlayerTankCrash),
	consts.SoundEnemyTankCrash:  newSound(soundEnemyTankCrash),
	consts.SoundBrickWallCrash:  newSound(soundBrickWallCrash),
	consts.SoundTankShoot:       newSound(soundTankShoot),
}

func newSound(soundFile string) *sound {
	s := sound{
		soundFile: soundFile,
	}
	s.Init()
	return &s
}

func (s *sound) Init() {
	f, err := os.Open(path.Join(soundPath, s.soundFile))
	if err != nil {
		panic(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		panic(err)
	}
	s.streamer = &streamer
	s.format = &format
}

func (s *sound) Play() {
	speaker.Init(s.format.SampleRate, s.format.SampleRate.N(time.Second/10))
	speaker.Play(*s.streamer)
}

func (s *sound) Close() {
	(*s.streamer).Close()
}

func (s *resourceManager) PlaySound(name consts.SoundType) {
	sound, ok := soundsMap[name]
	if !ok {
		panic(errors.New("Unable to find sound by name"))
	}
	sound.Play()
}

func (s *resourceManager) CloseSound() {
	for _, sound := range soundsMap {
		sound.Close()
	}
}
