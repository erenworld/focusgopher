package hosts

import (
	"reflect"
	"testing"
)

func TestExtractDomainsFromData(t *testing.T) {
	tests := []struct {
		name 		string
		data 		string
		wantDomains	[]string
		wantStatus  FocusStatus
		wantErr		bool
	}{
		{
			name:		 "no domain",
			data:		 "",
			wantDomains: []string{},
			wantStatus:	 FocusStatusOff,
			wantErr: 	 false,
		},
		{
			name:		 "no focusgopher domains",
			data:		 "",
			wantDomains: []string{},
			wantStatus:	 FocusStatusOff,
			wantErr: 	 false,
		},
		{
			name:		 "one domain and status on",
			data:		 `#focusgopher:start
							#focusgopher:on
							127.0.0.1 example.com
							127.0.0.1 example.com
							#focusgopher:end`,
			wantDomains: []string{"example.com"},
			wantStatus:	 FocusStatusOn,
			wantErr: 	 false,
		},
		{
			name:		 "no domain and status off",
			data:		 `#focusgopher:start
						  #focusgopher:end`,
			wantDomains: []string{},
			wantStatus:	 FocusStatusOff,
			wantErr: 	 false,
		},
		{
			name:		 "commented domains",
			data:		 `#focusgopher:start
						  # 127.0.0.1 example.com
						  #focusgopher:end`,
			wantDomains: []string{"example.com"},
			wantStatus:	 FocusStatusOff,
			wantErr: 	 false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDomains, gotStatus, err := ExtractDomainsFromData(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractDomainsFromData() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotDomains, tt.wantDomains) {
				t.Errorf("ExtractDomainsFromData() gotDomains = %v, want %v", gotDomains, tt.wantDomains)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("ExtractDomainsFromData() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		}) 
	}
}