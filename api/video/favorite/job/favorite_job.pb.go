// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: video/favorite/job/favorite_job.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_video_favorite_job_favorite_job_proto protoreflect.FileDescriptor

var file_video_favorite_job_favorite_job_proto_rawDesc = []byte{
	0x0a, 0x25, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6a, 0x6f,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x32, 0x0a, 0x0a, 0x08, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_video_favorite_job_favorite_job_proto_goTypes = []interface{}{}
var file_video_favorite_job_favorite_job_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_video_favorite_job_favorite_job_proto_init() }
func file_video_favorite_job_favorite_job_proto_init() {
	if File_video_favorite_job_favorite_job_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_video_favorite_job_favorite_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_favorite_job_favorite_job_proto_goTypes,
		DependencyIndexes: file_video_favorite_job_favorite_job_proto_depIdxs,
	}.Build()
	File_video_favorite_job_favorite_job_proto = out.File
	file_video_favorite_job_favorite_job_proto_rawDesc = nil
	file_video_favorite_job_favorite_job_proto_goTypes = nil
	file_video_favorite_job_favorite_job_proto_depIdxs = nil
}
