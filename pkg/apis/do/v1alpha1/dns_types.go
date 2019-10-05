package v1alpha1

import (
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DNSSpec defines the desired state of DNS
// +k8s:openapi-gen=true
type DNSSpec struct {
	DomainName string         `json:"domainName"`
	RecordType string         `json:"recordType"`
	Hostname   string         `json:"hostname"`
	Value      DNSRecordValue `json:"value"`
	TTL        *int           `json:"ttl,omitempty"`
	Port       *int           `json:"port,omitempty"`
	Priority   *int           `json:"priority,omitempty"`
	Flag       *int           `json:"flag,omitempty"`
	Weight     *int           `json:"weight,omitempty"`
	Tag        *string        `json:"tag,omitempty"`
}

type DNSRecordValue struct {
	Ref struct {
		IngressName string `json:"ingressName"`
	}
	Literal string `json:"literal"`
}

type RecordState int

func (s RecordState) values() [4]string {
	return [...]string{"initial", "pending", "active", "deleting"}
}

func (s RecordState) String() string {
	return s.values()[s]
}
func (s RecordState) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s *RecordState) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	for i, v := range s.values() {
		if v == str {
			*s = RecordState(i)
			return nil
		}
	}
	return nil
}

const (
	STATE_INITIAL RecordState = iota
	STATE_PENDING
	STATE_ACTIVE
	STATE_DELETING
)

// DNSStatus defines the observed state of DNS
// +k8s:openapi-gen=true
type DNSStatus struct {
	State RecordState `json:"state"`
	ID    int         `json:"id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNS is the Schema for the dns API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +genclient:nonNamespaced
type DNS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSSpec   `json:"spec,omitempty"`
	Status DNSStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSList contains a list of DNS
type DNSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNS `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DNS{}, &DNSList{})
}
