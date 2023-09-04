/*

  Simple DirectMedia Layer
  Copyright (C) 1997-2023 Sam Lantinga <slouken@libsdl.org>

  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.

  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:

  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.

*/

package sdl

import (
	"runtime.link/std"
)

type Audio struct {
	location

	Init func(AudioDriver) std.Error `ffi:"SDL_AudioInit"` // Internally used to initialize the audio subsystem with a specific driver.
	Quit func() std.Error            `ffi:"SDL_AudioQuit"` // Internally used to shut down the audio subsystem.

	DriverCount func() std.Int            `ffi:"SDL_GetNumAudioDrivers"`    // Get the number of built-in audio drivers.
	DriverIndex func(std.Int) AudioDriver `ffi:"SDL_GetAudioDriver"`        // Get the name of a built-in audio driver.
	Driver      func() AudioDriver        `ffi:"SDL_GetCurrentAudioDriver"` // Get the name of the current audio driver.

	Open   func(*AudioSpec) (std.Error, AudioSpec) `ffi:"SDL_OpenAudio"`      // Open a specific audio device.
	Status func() AudioStatus                      `ffi:"SDL_GetAudioStatus"` // Get the current audio state.
	Pause  func(Bool)                              `ffi:"SDL_PauseAudio"`     // Pause and unpause the audio callback processing.

	LoadWAV func(src *File, free_source std.Int, spec *AudioSpec, buf *std.Buffer) `ffi:"SDL_LoadWAV_RW"` // Load a WAVE from an SDL_RWops object.
	FreeWAV func(std.Buffer)                                                       `ffi:"SDL_FreeWAV"`    // Free an audio buffer previously allocated with LoadWAV().

	BuildCVT func(cvt *AudioCVT, fmt AudioFormat, src_channels std.Uint8, src_rate std.Int, dst_format AudioFormat, dst_channels std.Uint8, dst_rate std.Int) std.Error `ffi:"SDL_BuildAudioCVT"` // Initialize a CVT structure for conversion.
	Convert  func(cvt *AudioCVT) std.Error                                                                                                                              `ffi:"SDL_ConvertAudio"`  // Convert audio data to a desired audio format.

	Mix       func(dst, src *std.Uint8, len std.Uint32, volume std.Int)                     `ffi:"SDL_MixAudio"`       // Mix audio data in a specified format.
	MixFormat func(dst, src *std.Uint8, format AudioFormat, len std.Uint32, volume std.Int) `ffi:"SDL_MixAudioFormat"` // Mix audio data in a specified format.

	Lock   func() `ffi:"SDL_LockAudio"`   // Lock out the callback function.
	Unlock func() `ffi:"SDL_UnlockAudio"` // Unlock the callback function.
	Close  func() `ffi:"SDL_CloseAudio"`  // Close the audio device previously opened with Open().
}

type AudioStreams struct {
	location

	New func(src_format AudioFormat, src_channels std.Uint8, src_rate std.Int, dst_format AudioFormat, dst_channels std.Uint8, dst_rate std.Int) (AudioStream, error) `ffi:"SDL_NewAudioStream"` // Create a new audio stream.

	Put   func(stream AudioStream, buf std.Buffer) std.Error `ffi:"SDL_AudioStreamPut"`   // Write data to a stream.
	Get   func(stream AudioStream) std.Int                   `ffi:"SDL_AudioStreamGet"`   // Read data from a stream.
	Flush func(stream AudioStream) std.Int                   `ffi:"SDL_AudioStreamFlush"` // Flush any pending data in the stream.
	Clear func(stream AudioStream)                           `ffi:"SDL_AudioStreamClear"` // Clear any pending data in the stream, without flushing.
	Free  func(stream AudioStream)                           `ffi:"SDL_FreeAudioStream"`  // Free an audio stream.
}

type AudioDevices struct {
	location

	Default func(*std.String, *AudioSpec, std.Int) std.Error `ffi:"SDL_GetDefaultAudioInfo"` // Get the ID of a built-in audio device that is the "best" fit for the desired device specification.

	Open   func(AudioDeviceName, std.Int, *AudioSpec, *AudioSpec, AudioAllowedChanges) (std.Error, AudioDevice) `ffi:"SDL_OpenAudioDevice"`      // Open a specific audio device.
	Count  func(std.Int) AudioDeviceIndex                                                                       `ffi:"SDL_GetNumAudioDevices"`   // Get the number of available devices exposed by the current driver.
	Name   func(AudioDeviceIndex, std.Int) AudioDeviceName                                                      `ffi:"SDL_GetAudioDeviceName"`   // Get the human-readable name of a specific audio device.
	Spec   func(AudioDeviceIndex, std.Int) (std.Error, AudioSpec)                                               `ffi:"SDL_GetAudioDeviceSpec"`   // Get the audio device specification for a specific device.
	Pause  func(AudioDevice, Bool)                                                                              `ffi:"SDL_PauseAudioDevice"`     // Pause and unpause a specific audio device.
	Status func(AudioDevice) AudioStatus                                                                        `ffi:"SDL_GetAudioDeviceStatus"` // Get the current audio state of a specific device.
	Lock   func(AudioDevice)                                                                                    `ffi:"SDL_LockAudioDevice"`      // Lock the audio device mutex.
	Unlock func(AudioDevice)                                                                                    `ffi:"SDL_UnlockAudioDevice"`    // Unlock the audio device mutex.
	Close  func(AudioDevice)                                                                                    `ffi:"SDL_CloseAudioDevice"`     // Close a specific audio device.

	Queue      func(device AudioDevice, data std.UnsafePointer, len std.Uint32) std.Error `ffi:"SDL_QueueAudio"`         // Queue more audio to playback on a specific device.
	Dequeue    func(device AudioDevice, data std.UnsafePointer, len std.Uint32) std.Error `ffi:"SDL_DequeueAudio"`       // Dequeue more audio for playback on a specific device.
	QueuedSuze func(device AudioDevice) std.Uint32                                        `ffi:"SDL_GetQueuedAudioSize"` // Get the number of bytes of still-queued audio.
	ClearQueue func(device AudioDevice)                                                   `ffi:"SDL_ClearQueuedAudio"`   // Drop any queued audio data.
}

type AudioDriver string

type AudioStream std.Handle[AudioStream]

type AudioDevice std.Uint32
type AudioDeviceIndex std.Int
type AudioDeviceName string

type AudioStatus std.Enum

const (
	AudioStopped AudioStatus = iota
	AudioPlaying
	AudioPaused
)

type AudioFormat std.Uint16

const (
	AudioFormatMaskBitSize  AudioFormat = 0xFF
	AudioFormatMaskDataType AudioFormat = 1 << 8
	AudioFormatMaskEndian   AudioFormat = 1 << 12
	AudioFormatMaskSigned   AudioFormat = 1 << 15
)

func (af AudioFormat) BitSize() std.Uint16 {
	return std.Uint16(af & AudioFormatMaskBitSize)
}

func (af AudioFormat) IsFloat() bool {
	return af&AudioFormatMaskDataType != 0
}

func (af AudioFormat) IsBigEndian() bool {
	return af&AudioFormatMaskEndian != 0
}

func (af AudioFormat) IsSigned() bool {
	return af&AudioFormatMaskSigned != 0
}

func (af AudioFormat) IsInt() bool {
	return !af.IsFloat()
}

func (af AudioFormat) IsLittleEndian() bool {
	return !af.IsBigEndian()
}

func (af AudioFormat) IsUnsigned() bool {
	return !af.IsSigned()
}

const (
	AudioU8     AudioFormat = 0x0008 /**< Unsigned 8-bit samples */
	AudioS8     AudioFormat = 0x8008 /**< Signed 8-bit samples */
	AudioU16LSB AudioFormat = 0x0010 /**< Unsigned 16-bit samples */
	AudioS16LSB AudioFormat = 0x8010 /**< Signed 16-bit samples */
	AudioU16MSB AudioFormat = 0x1010 /**< As above, but big-endian byte order */
	AudioS16MSB AudioFormat = 0x9010 /**< As above, but big-endian byte order */
	AudioU16    AudioFormat = AudioU16LSB
	AudioS16    AudioFormat = AudioS16LSB
)

const (
	AudioS32LSB AudioFormat = 0x8020 /**< 32-bit integer samples */
	AudioS32MSB AudioFormat = 0x9020 /**< As above, but big-endian byte order */
	AudioS32    AudioFormat = AudioS32LSB
)

const (
	AudioF32LSB AudioFormat = 0x8120 /**< 32-bit floating point samples */
	AudioF32MSB AudioFormat = 0x9120 /**< As above, but big-endian byte order */
	AudioF32    AudioFormat = AudioF32LSB
)

type AudioAllowedChanges std.Int

const (
	AudioAllowFrequencyChange AudioAllowedChanges = 0x00000001 /**< Allow any sample rate for playback */
	AudioAllowFormatChange    AudioAllowedChanges = 0x00000002 /**< Allow any audio format for playback */
	AudioAllowChannelsChange  AudioAllowedChanges = 0x00000004 /**< Allow any number of channels for playback */
	AudioAllowSamplesChange   AudioAllowedChanges = 0x00000008 /**< Allow any number of samples for playback */
	AudioAllowAnyChange       AudioAllowedChanges = AudioAllowFrequencyChange | AudioAllowFormatChange | AudioAllowChannelsChange | AudioAllowSamplesChange
)

type AudioCallback std.Func[func(std.UnsafePointer, std.Buffer)]

type AudioSpec struct {
	Freq     std.Int           /**< DSP frequency -- samples per second */
	Format   AudioFormat       /**< Audio data format */
	Channels std.Uint8         /**< Number of channels: 1 mono, 2 stereo */
	Silence  std.Uint8         /**< Audio buffer silence value (calculated) */
	Samples  std.Uint16        /**< Audio buffer size in samples (power of 2) */
	Padding  std.Uint16        /**< Necessary for some compile environments */
	Size     std.Uint32        /**< Audio buffer size in bytes (calculated) */
	Callback AudioCallback     /**< Callback that feeds the audio device (NULL to use SDL_QueueAudio()). */
	Userdata std.UnsafePointer /**< Userdata passed to callback (ignored for NULL callbacks). */
}

type AudioFilter std.Func[func(*AudioCVT, AudioFormat)]

const MaxFiltersAudioCVT = 9

type AudioCVT struct {
	Needed                  std.Int                             /**< Set to 1 if conversion possible */
	SourceFormat            AudioFormat                         /**< Source audio format */
	TargetFormat            AudioFormat                         /**< Target audio format */
	RateConversionIncrement std.Double                          /**< Rate conversion increment */
	Buffer                  std.Buffer                          /**< Buffer to hold entire audio data */
	LengthConverted         std.Int                             /**< Length of converted audio buffer */
	LengthMultiple          std.Int                             /**< buffer must be len*mul in length */
	LengthRatio             std.Double                          /**< Given len, final size is len*rat */
	Filters                 [MaxFiltersAudioCVT + 1]AudioFilter /**< NULL-terminated list of Filter functions */
	FilterIndex             std.Int                             /**< Current audio conversion function */
}
