package sarama

import "testing"

var (
	offsetRequestNoBlocksV1 = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0x00,
	}

	offsetRequestNoBlocksV2 = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0x00,
		0x00,
	}

	offsetRequestOneBlock = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x03, 'f', 'o', 'o',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x04,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
	}

	offsetRequestOneBlockV1 = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x03, 'b', 'a', 'r',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x04,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
	}

	offsetRequestOneBlockReadCommittedV2 = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0x01, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x03, 'b', 'a', 'r',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x04,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
	}

	offsetRequestReplicaID = []byte{
		0x00, 0x00, 0x00, 0x2a,
		0x00, 0x00, 0x00, 0x00,
	}
)

func TestOffsetRequest(t *testing.T) {
	t.Parallel()
	request := new(OffsetRequest)
	testRequest(t, "no blocks", request, offsetRequestNoBlocksV1)

	request.AddBlock("foo", 4, 1, 2)
	testRequest(t, "one block", request, offsetRequestOneBlock)
}

func TestOffsetRequestV1(t *testing.T) {
	t.Parallel()
	request := new(OffsetRequest)
	request.Version = 1
	testRequest(t, "no blocks", request, offsetRequestNoBlocksV1)

	request.AddBlock("bar", 4, 1, 2) // Last argument is ignored for V1
	testRequest(t, "one block", request, offsetRequestOneBlockV1)
}

func TestOffsetRequestV2(t *testing.T) {
	t.Parallel()
	request := new(OffsetRequest)
	request.Version = 2
	testRequest(t, "no blocks", request, offsetRequestNoBlocksV2)

	request.IsolationLevel = ReadCommitted
	request.AddBlock("bar", 4, 1, 2) // Last argument is ignored for V1
	testRequest(t, "one block", request, offsetRequestOneBlockReadCommittedV2)
}

func TestOffsetRequestReplicaID(t *testing.T) {
	t.Parallel()
	request := new(OffsetRequest)
	replicaID := int32(42)
	request.SetReplicaID(replicaID)

	if found := request.ReplicaID(); found != replicaID {
		t.Errorf("replicaID: expected %v, found %v", replicaID, found)
	}

	testRequest(t, "with replica ID", request, offsetRequestReplicaID)
}
