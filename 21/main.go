package main

import (
	"fmt"
)

// MediaPlayer - интерфейс, который определяет метод Play для воспроизведения медиафайлов.
type MediaPlayer interface {
	Play(audioType string, fileName string) error
}

// AdvancedMediaPlayer - интерфейс, который определяет методы PlayVideo и PlayAudio для воспроизведения видео и аудио соответственно.
type AdvancedMediaPlayer interface {
	PlayVideo(fileName string) error
	PlayAudio(fileName string) error
}

// VLCPlayer - класс, реализующий интерфейс AdvancedMediaPlayer для воспроизведения видео и аудио.
type VLCPlayer struct{}

// PlayVideo - метод для воспроизведения видеофайла.
func (v *VLCPlayer) PlayVideo(fileName string) error {
	fmt.Printf("Playing video file. Name: %s\n", fileName)
	return nil
}

// PlayAudio - метод для воспроизведения аудиофайла.
func (v *VLCPlayer) PlayAudio(fileName string) error {
	fmt.Printf("Playing audio file. Name: %s\n", fileName)
	return nil
}

// MediaAdapter - адаптер, который адаптирует интерфейс AdvancedMediaPlayer к интерфейсу MediaPlayer.
type MediaAdapter struct {
	advancedMusicPlayer AdvancedMediaPlayer
}

// Play - метод адаптера, который реализует метод Play интерфейса MediaPlayer.
func (m *MediaAdapter) Play(audioType string, fileName string) error {
	if audioType == "vlc" {
		return m.advancedMusicPlayer.PlayVideo(fileName)
	} else if audioType == "mp4" {
		return m.advancedMusicPlayer.PlayAudio(fileName)
	}
	return fmt.Errorf("invalid media. %s format not supported", audioType)
}

// AudioPlayer - класс, который использует адаптер MediaAdapter для воспроизведения медиафайлов.
type AudioPlayer struct {
	mediaAdapter MediaPlayer
}

// Play - метод класса AudioPlayer для воспроизведения медиафайлов.
func (a *AudioPlayer) Play(audioType string, fileName string) error {
	// Встроенная поддержка mp3 музыкальных файлов
	if audioType == "mp3" {
		fmt.Printf("Playing mp3 file. Name: %s\n", fileName)
		return nil
	} else if audioType == "vlc" || audioType == "mp4" {

		// MediaAdapter предоставляет поддержку для других форматов файлов
		a.mediaAdapter = &MediaAdapter{&VLCPlayer{}}
		return a.mediaAdapter.Play(audioType, fileName)
	}

	return fmt.Errorf("invalid media. %s format not supported", audioType)
}

func main() {
	// Создаем экземпляр AudioPlayer
	audioPlayer := AudioPlayer{}

	// Воспроизводим mp3 файл
	err := audioPlayer.Play("mp3", "example.mp3")
	if err != nil {
		fmt.Println(err)
	}

	// Воспроизводим vlc файл
	err = audioPlayer.Play("vlc", "example.vlc")
	if err != nil {
		fmt.Println(err)
	}

	// Воспроизводим mp4 файл
	err = audioPlayer.Play("mp4", "example.mp4")
	if err != nil {
		fmt.Println(err)
	}
}
