// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElastictranscoderPreset struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElastictranscoderPresetSpec `json"spec"`
}

type AwsElastictranscoderPresetSpec struct {
	Video             AwsElastictranscoderPresetSpecVideo             `json:"video"`
	VideoWatermarks   AwsElastictranscoderPresetSpecVideoWatermarks   `json:"video_watermarks"`
	VideoCodecOptions map[string]string                               `json:"video_codec_options"`
	Arn               string                                          `json:"arn"`
	Description       string                                          `json:"description"`
	Name              string                                          `json:"name"`
	Thumbnails        AwsElastictranscoderPresetSpecThumbnails        `json:"thumbnails"`
	Type              string                                          `json:"type"`
	Audio             AwsElastictranscoderPresetSpecAudio             `json:"audio"`
	AudioCodecOptions AwsElastictranscoderPresetSpecAudioCodecOptions `json:"audio_codec_options"`
	Container         string                                          `json:"container"`
}

type AwsElastictranscoderPresetSpecVideo struct {
	Resolution         string `json:"resolution"`
	AspectRatio        string `json:"aspect_ratio"`
	DisplayAspectRatio string `json:"display_aspect_ratio"`
	FrameRate          string `json:"frame_rate"`
	PaddingPolicy      string `json:"padding_policy"`
	MaxFrameRate       string `json:"max_frame_rate"`
	MaxHeight          string `json:"max_height"`
	MaxWidth           string `json:"max_width"`
	SizingPolicy       string `json:"sizing_policy"`
	BitRate            string `json:"bit_rate"`
	Codec              string `json:"codec"`
	FixedGop           string `json:"fixed_gop"`
	KeyframesMaxDist   string `json:"keyframes_max_dist"`
}

type AwsElastictranscoderPresetSpecVideoWatermarks struct {
	HorizontalOffset string `json:"horizontal_offset"`
	Opacity          string `json:"opacity"`
	SizingPolicy     string `json:"sizing_policy"`
	Target           string `json:"target"`
	VerticalAlign    string `json:"vertical_align"`
	HorizontalAlign  string `json:"horizontal_align"`
	MaxHeight        string `json:"max_height"`
	MaxWidth         string `json:"max_width"`
	VerticalOffset   string `json:"vertical_offset"`
	Id               string `json:"id"`
}

type AwsElastictranscoderPresetSpecThumbnails struct {
	AspectRatio   string `json:"aspect_ratio"`
	Format        string `json:"format"`
	Interval      string `json:"interval"`
	MaxHeight     string `json:"max_height"`
	MaxWidth      string `json:"max_width"`
	PaddingPolicy string `json:"padding_policy"`
	Resolution    string `json:"resolution"`
	SizingPolicy  string `json:"sizing_policy"`
}

type AwsElastictranscoderPresetSpecAudio struct {
	Codec            string `json:"codec"`
	SampleRate       string `json:"sample_rate"`
	AudioPackingMode string `json:"audio_packing_mode"`
	BitRate          string `json:"bit_rate"`
	Channels         string `json:"channels"`
}

type AwsElastictranscoderPresetSpecAudioCodecOptions struct {
	BitDepth string `json:"bit_depth"`
	BitOrder string `json:"bit_order"`
	Profile  string `json:"profile"`
	Signed   string `json:"signed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElastictranscoderPresetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElastictranscoderPreset `json"items"`
}
