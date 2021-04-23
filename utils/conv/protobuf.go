package conv

import (
	pb "k3s-nclink-apps/configmodel"
	"k3s-nclink-apps/data-source/entity"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func MessageToString(msg proto.Message) (string, error) {
	json, err := protojson.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(json), err
}

func DbModelToWireModel(in *entity.Model) (*pb.Model, error) {
	ret := &pb.Model{}
	err := protojson.Unmarshal([]byte(in.Def), ret)
	if err != nil {
		return nil, err
	}
	ret.Id = in.Id
	return ret, nil
}
