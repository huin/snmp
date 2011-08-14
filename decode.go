package snmp

import (
	"asn1"
	"fmt"
	"os"
)

func decode(data []byte) (interface{}, os.Error) {
        m := Message{}
        _, err := asn1.Unmarshal(data, &m)
        if err != nil {
		fmt.Errorf("%#v", data)
                return nil, err
        }
	switch m.Data.FullBytes[0] {
	case 0xa0:
		// GetRequest
		request := new(GetRequest)
		// hack ANY -> IMPLICIT SEQUENCE
                m.Data.FullBytes[0] = 0x30
                _, err = asn1.Unmarshal(m.Data.FullBytes, &request)
                if err != nil {
                        return nil, fmt.Errorf("%#v, %#v, %s",m.Data.FullBytes, request, err)
                }
		return request, nil
	case 0xa1:
		// GetNextRequest
	case 0xa2:
		// Response
		response := Response{}
		// hack ANY -> IMPLICIT SEQUENCE
        	m.Data.FullBytes[0] = 0x30
        	_, err = asn1.Unmarshal(m.Data.FullBytes, &response)
        	if err != nil {
                	return nil, fmt.Errorf("%#v, %#v, %s",m.Data.FullBytes, response, err)
        	}
		return &response, nil
	case 0xa3:
		// SetResponse
	case 0xa4:
		// InformRequest
	default:
	}
	return nil, fmt.Errorf("Unknown CHOICE")
}