package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dh1tw/remoteAudio/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "remoteAudio",
	Short: "Audio streaming client & server for remote Amateur radio operations",
	Long:  `Audio streaming client & server for remote Amateur radio operations`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.remoteAudio.yaml)")
	RootCmd.PersistentFlags().StringP("input-device-name", "i", "default", "Input device")
	RootCmd.PersistentFlags().Float64("input-device-samplingrate", 48000, "Input device sampling rate")
	RootCmd.PersistentFlags().Duration("input-device-latency", time.Millisecond*5, "Input latency")
	RootCmd.PersistentFlags().String("input-device-channels", "mono", "Input Channels")

	RootCmd.PersistentFlags().StringP("output-device-name", "o", "default", "Output device")
	RootCmd.PersistentFlags().Float64("output-device-samplingrate", 48000, "Output device sampling rate")
	RootCmd.PersistentFlags().Duration("output-device-latency", time.Millisecond*5, "Output latency")
	RootCmd.PersistentFlags().String("output-device-channels", "stereo", "Output Channels")

	RootCmd.PersistentFlags().Float64("pcm-samplingrate", 16000, "pcm sampling rate")
	RootCmd.PersistentFlags().Int("pcm-bitdepth", 16, "pcm audio bit depth (8, 12, 16, 24 bit)")
	RootCmd.PersistentFlags().String("pcm-channels", "stereo", "pcm audio Channels")
	RootCmd.PersistentFlags().Int("pcm-resampling-quality", 1, "pcm resampling quality")

	RootCmd.PersistentFlags().String("opus-application", "restricted_lowdelay", "profile for opus encoder")
	RootCmd.PersistentFlags().Int("opus-bitrate", 32000, "Bitrate (bits/sec) generated by the opus encoder")
	RootCmd.PersistentFlags().Int("opus-complexity", 9, "Computational complexity of opus encoder")
	RootCmd.PersistentFlags().String("opus-max-bandwidth", "wideband", "maximum bandwidth of opus encoder")

	RootCmd.PersistentFlags().IntP("audio-frame-length", "f", 960, "Amount of audio samples in one frame")
	RootCmd.PersistentFlags().IntP("rx-buffer-length", "R", 10, "Buffer length (in frames) for incoming Audio packets")
	RootCmd.PersistentFlags().StringP("codec", "c", "opus", "Audio codec")

	// hidden flags
	// RootCmd.PersistentFlags().BoolVar(&profServerEnabled, "prof-server", false, "enable profiling server at http://0.0.0.0:6060/debug/pprof")
	// RootCmd.PersistentFlags().MarkHidden("prof-server")

	viper.BindPFlag("input-device.device-name", RootCmd.PersistentFlags().Lookup("input-device-name"))
	viper.BindPFlag("input-device.samplingrate", RootCmd.PersistentFlags().Lookup("input-device-samplingrate"))
	viper.BindPFlag("input-device.latency", RootCmd.PersistentFlags().Lookup("input-device-latency"))
	viper.BindPFlag("input-device.channels", RootCmd.PersistentFlags().Lookup("input-device-channels"))

	viper.BindPFlag("output-device.device-name", RootCmd.PersistentFlags().Lookup("output-device-name"))
	viper.BindPFlag("output-device.samplingrate", RootCmd.PersistentFlags().Lookup("output-device-samplingrate"))
	viper.BindPFlag("output-device.latency", RootCmd.PersistentFlags().Lookup("output-device-latency"))
	viper.BindPFlag("output-device.channels", RootCmd.PersistentFlags().Lookup("output-device-channels"))

	viper.BindPFlag("pcm.samplingrate", RootCmd.PersistentFlags().Lookup("pcm-samplingrate"))
	viper.BindPFlag("pcm.bitdepth", RootCmd.PersistentFlags().Lookup("pcm-bitdepth"))
	viper.BindPFlag("pcm.channels", RootCmd.PersistentFlags().Lookup("pcm-channels"))
	viper.BindPFlag("pcm.resampling-quality", RootCmd.PersistentFlags().Lookup("pcm-resampling-quality"))

	viper.BindPFlag("opus.application", RootCmd.PersistentFlags().Lookup("opus-application"))
	viper.BindPFlag("opus.bitrate", RootCmd.PersistentFlags().Lookup("opus-bitrate"))
	viper.BindPFlag("opus.complexity", RootCmd.PersistentFlags().Lookup("opus-complexity"))
	viper.BindPFlag("opus.max-bandwidth", RootCmd.PersistentFlags().Lookup("opus-max-bandwidth"))

	viper.BindPFlag("audio.frame-length", RootCmd.PersistentFlags().Lookup("audio-frame-length"))
	viper.BindPFlag("audio.rx-buffer-length", RootCmd.PersistentFlags().Lookup("rx-buffer-length"))
	viper.BindPFlag("audio.codec", RootCmd.PersistentFlags().Lookup("codec"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".remoteAudio") // name of config file (without extension)
		viper.AddConfigPath("$HOME")        // adding home directory as first search path
	}

	viper.AutomaticEnv() // read in environment variables that match
}

func checkAudioParameterValues() bool {

	ok := true

	if chs := strings.ToUpper(viper.GetString("input-device.channels")); chs != "MONO" && chs != "STEREO" {
		fmt.Println(&parmError{"input-device.channels", "allowed values are [MONO, STEREO]"})
		ok = false
	}

	if chs := strings.ToUpper(viper.GetString("output-device.channels")); chs != "MONO" && chs != "STEREO" {
		fmt.Println(&parmError{"output-device.channels", "allowed values are [MONO, STEREO]"})
		ok = false
	}

	if codec := strings.ToUpper(viper.GetString("audio.codec")); codec != "OPUS" && codec != "PCM" {
		fmt.Println(&parmError{"audio.codec", "allowed values are [OPUS, PCM]"})
		ok = false
	}

	if strings.ToUpper(viper.GetString("audio.codec")) == "PCM" {

		if viper.GetFloat64("pcm.samplingrate") < 0 {
			fmt.Println(&parmError{"pcm.samplingrate", "value must be > 0"})
			ok = false
		}

		if viper.GetFloat64("pcm.samplingrate")/viper.GetFloat64("input-device.samplingrate") < 1/256 ||
			viper.GetFloat64("pcm.samplingrate")/viper.GetFloat64("input-device.samplingrate") > 256 ||
			viper.GetFloat64("input-device.samplingrate")/viper.GetFloat64("pcm.samplingrate") < 1/256 ||
			viper.GetFloat64("input-device.samplingrate")/viper.GetFloat64("pcm.samplingrate") > 256 {
			fmt.Println(&parmError{"pcm.samplingrate", "ratio between input-device & pcm samplingrate must be < 256"})
			ok = false
		}

		if viper.GetInt("pcm.bitdepth") != 8 &&
			viper.GetInt("pcm.bitdepth") != 12 &&
			viper.GetInt("pcm.bitdepth") != 16 &&
			viper.GetInt("pcm.bitdepth") != 24 {
			fmt.Println(&parmError{"pcm.bitdepth", "allowed values are [8, 12, 16, 24]"})
			ok = false
		}

		if chs := strings.ToUpper(viper.GetString("pcm.channels")); chs != "MONO" && chs != "STEREO" {
			fmt.Println(&parmError{"pcm.channels", "allowed values are [MONO, STEREO]"})
			ok = false
		}

		if viper.GetInt("pcm.resampling-quality") < 0 || viper.GetInt("pcm.resampling-quality") > 4 {
			fmt.Println(&parmError{"pcm.resampling-quality", "allowed values are [0...4]"})
			ok = false
		}

		if viper.GetInt("audio.frame-length") <= 0 {
			fmt.Println(&parmError{"audio.frame-length", "value must be > 0"})
			ok = false
		}
	}

	if strings.ToUpper(viper.GetString("audio.codec")) == "OPUS" {
		opusApps := []string{"RESTRICTED_LOWDELAY", "VOIP", "AUDIO"}
		opusApp := strings.ToUpper(viper.GetString("opus.application"))
		if !utils.StringInSlice(opusApp, opusApps) {
			fmt.Println(&parmError{"opus.application", "allowed values are VOIP, AUDIO or RESTRICTED_LOWDELAY"})
			ok = false
		}

		opusMaxBws := []string{"NARROWBAND", "MEDIUMBAND", "WIDEBAND", "SUPERWIDEBAND", "FULLBAND"}
		opusBw := strings.ToUpper(viper.GetString("opus.max-bandwidth"))
		if !utils.StringInSlice(opusBw, opusMaxBws) {
			fmt.Println(&parmError{"opus.max-bandwidth", "allowed values are NARROWBAND, MEDIUMBAND, WIDEBAND, SUPERWIDEBAND, FULLBAND"})
			ok = false
		}

		if viper.GetInt("opus.bitrate") < 6000 || viper.GetInt("opus.bitrate") > 510000 {
			fmt.Println(&parmError{"opus.bitrate", "allowed values are [6000...510000]"})
			ok = false
		}

		if viper.GetInt("opus.complexity") < 0 || viper.GetInt("opus.complexity") > 10 {
			fmt.Println(&parmError{"opus.complexity", "allowed values are [0...10]"})
			ok = false
		}

		opusFrameLength := float64(viper.GetInt("audio.frame-length")) / viper.GetFloat64("input-device.samplingrate")
		if opusFrameLength != 0.0025 &&
			opusFrameLength != 0.005 &&
			opusFrameLength != 0.01 &&
			opusFrameLength != 0.02 &&
			opusFrameLength != 0.04 &&
			opusFrameLength != 0.06 {
			fmt.Println(&parmError{"audio.frame-length", "division of audio.frame-length/input-device.samplingrate must result in 2.5, 5, 10, 20, 40, 60ms for the opus codec"})
			ok = false
		}
	}

	if viper.GetInt("audio.rx-buffer-length") <= 0 {
		fmt.Println(&parmError{"audio.rx-buffer-length", "value must be > 0"})
		ok = false
	}

	return ok
}
