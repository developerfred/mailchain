package nacl

import (
	"testing"

	"github.com/mailchain/mailchain/crypto"
	"github.com/mailchain/mailchain/crypto/cipher"
	"github.com/mailchain/mailchain/crypto/ed25519/ed25519test"
	"github.com/mailchain/mailchain/crypto/secp256k1/secp256k1test"
	"github.com/mailchain/mailchain/crypto/sr25519/sr25519test"
	"github.com/stretchr/testify/assert"
)

func TestNewDecrypter(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		privateKey crypto.PrivateKey
	}
	tests := []struct {
		name    string
		args    args
		want    *Decrypter
		wantErr bool
	}{
		{
			"success-charlotte-ed25519",
			args{
				ed25519test.CharlottePrivateKey,
			},
			&Decrypter{
				privateKey: ed25519test.CharlottePrivateKey,
			},
			false,
		},
		{
			"success-sofia-ed25519",
			args{
				ed25519test.SofiaPrivateKey,
			},
			&Decrypter{
				privateKey: ed25519test.SofiaPrivateKey,
			},
			false,
		},
		{
			"success-sofia-sr25519",
			args{
				sr25519test.SofiaPrivateKey,
			},
			&Decrypter{
				privateKey: sr25519test.SofiaPrivateKey,
			},
			false,
		},
		{
			"success-charlotte-sr25519",
			args{
				sr25519test.CharlottePrivateKey,
			},
			&Decrypter{
				privateKey: sr25519test.CharlottePrivateKey,
			},
			false,
		},
		{
			"invalid-key",
			args{
				secp256k1test.CharlottePrivateKey,
			},
			nil,
			true,
		},
		{
			"invalid-key-schnorrkel",
			args{
				secp256k1test.CharlottePrivateKey,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDecrypter(tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDecrypter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(tt.want, got) {
				t.Errorf("NewDecrypter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypter_Decrypt(t *testing.T) {
	assert := assert.New(t)
	type fields struct {
		privateKey crypto.PrivateKey
	}
	type args struct {
		data cipher.EncryptedContent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    cipher.PlainContent
		wantErr bool
	}{
		{
			"success-charlotte-ed25519",
			fields{
				ed25519test.CharlottePrivateKey,
			},
			args{
				cipher.EncryptedContent{0x2a, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x5b, 0x19, 0x83, 0xe5, 0x6e, 0x7f, 0xed, 0xfe, 0xbb, 0xd0, 0x70, 0x34, 0xce, 0x25, 0x49, 0x76, 0xa3, 0x50, 0x78, 0x91, 0x18, 0xe6, 0xe3},
			},
			cipher.PlainContent{0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65},
			false,
		},
		{
			"err-invalid-prefix",
			fields{
				ed25519test.CharlottePrivateKey,
			},
			args{
				cipher.EncryptedContent{0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x5b, 0x19, 0x83, 0xe5, 0x6e, 0x7f, 0xed, 0xfe, 0xbb, 0xd0, 0x70, 0x34, 0xce, 0x25, 0x49, 0x76, 0xa3, 0x50, 0x78, 0x91, 0x18, 0xe6, 0xe3},
			},
			nil,
			true,
		},
		{
			"err-wrong-key",
			fields{
				ed25519test.SofiaPrivateKey,
			},
			args{
				cipher.EncryptedContent{0x2a, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x5b, 0x19, 0x83, 0xe5, 0x6e, 0x7f, 0xed, 0xfe, 0xbb, 0xd0, 0x70, 0x34, 0xce, 0x25, 0x49, 0x76, 0xa3, 0x50, 0x78, 0x91, 0x18, 0xe6, 0xe3},
			},
			nil,
			true,
		},
		{
			"err-charlotte-secp256k1",
			fields{
				secp256k1test.CharlottePrivateKey,
			},
			args{
				cipher.EncryptedContent{0x2a, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x5b, 0x19, 0x83, 0xe5, 0x6e, 0x7f, 0xed, 0xfe, 0xbb, 0xd0, 0x70, 0x34, 0xce, 0x25, 0x49, 0x76, 0xa3, 0x50, 0x78, 0x91, 0x18, 0xe6, 0xe3},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Decrypter{
				privateKey: tt.fields.privateKey,
			}
			got, err := d.Decrypt(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypter.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(tt.want, got) {
				t.Errorf("Decrypter.Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validatePrivateKeyType(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		pk crypto.PrivateKey
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"success-ed25519-sofia",
			args{
				ed25519test.SofiaPrivateKey,
			},
			[]byte{0x72, 0x3c, 0xaa, 0x23, 0xa5, 0xb5, 0x11, 0xaf, 0x5a, 0xd7, 0xb7, 0xef, 0x60, 0x76, 0xe4, 0x14, 0xab, 0x7e, 0x75, 0xa9, 0xdc, 0x91, 0xe, 0xa6, 0xe, 0x41, 0x7a, 0x2b, 0x77, 0xa, 0x56, 0x71},
			false,
		},
		{
			"success-sr25519-sofia",
			args{
				sr25519test.SofiaPrivateKey,
			},
			[]byte{0x5c, 0x6d, 0x7a, 0xdf, 0x75, 0xbd, 0xa1, 0x18, 0x0c, 0x22, 0x5d, 0x25, 0xf3, 0xaa, 0x8d, 0xc1, 0x74, 0xbb, 0xfb, 0x3c, 0xdd, 0xee, 0x11, 0xae, 0x9a, 0x85, 0x98, 0x2f, 0x6f, 0xaf, 0x79, 0x1a},
			false,
		},
		{
			"success-sr25519-charlotte",
			args{
				sr25519test.CharlottePrivateKey,
			},
			[]byte{0x23, 0xb0, 0x63, 0xa5, 0x81, 0xfd, 0x8e, 0x5e, 0x84, 0x7c, 0x4e, 0x2b, 0x9c, 0x49, 0x42, 0x47, 0x29, 0x87, 0x91, 0x53, 0x0f, 0x52, 0x93, 0xbe, 0x36, 0x9e, 0x8b, 0xf2, 0x3a, 0x45, 0xd2, 0xbd},
			false,
		},
		{
			"err-secp256k1-sofia",
			args{
				secp256k1test.CharlottePrivateKey,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validatePrivateKeyType(tt.args.pk)
			if (err != nil) != tt.wantErr {
				t.Errorf("validatePrivateKeyType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(tt.want, got) {
				t.Errorf("validatePrivateKeyType() = %v, want %v", got, tt.want)
			}
		})
	}
}
