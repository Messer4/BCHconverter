package bchaddr

import (
	"github.com/Messer4/base58check"
	"errors"
	"github.com/thoas/go-funk"
	"bytes"
)

	type decod struct {
		hash []uint8
		format string
		network string
		tp string
	}

	var VERSION_BYTE = map[string]map[string]byte{
		"mainnet":{"P2PKH":0,"P2SH":5},
		"testnet":{"P2PKH":111,"P2SH":196},
		"regtest":{"P2PKH":111},
		}


/**
* Translates the given address into cashaddr format.
*/

	func ToCashAddress (address string) (string, error) {
		var prefix string
		//prefix:= "bchreg"
		d,err := decodeAddress(address)
		if d.network=="mainnet"{
			prefix = "bitcoincash"
		}else{
			prefix = "bchtest"
		}

		addr,err := encode(prefix,d.tp,d.hash)
		return addr,err
	}

	func ToCashAddressRGT (address string) (string, error) {
			prefix:= "bchreg"
		d,err := decodeAddress(address)
		addr,err := encode(prefix,d.tp,d.hash)
		return addr,err
	}

/**
* Translates the given address into legacy format.
*/
	func ToLegacyAddress (address string) (string,error) {
		decoded,err := decodeCashAddress(address)
	/*	if (decoded.format === Format.Legacy) {
			return address
		}*/
		return encodeAsLegacy(decoded),err
	}

/**
 * Encodes the given decoded address into legacy format.
 */
	func encodeAsLegacy (dec decod) (string) {
		var versionByte = VERSION_BYTE[dec.network][dec.tp]
		var buffer = bytes.Buffer{}
		//Buffer.alloc(1 + len(dec.hash))
		buffer.WriteByte(versionByte)
		buffer.Write(dec.hash)
		return base58check.Encode(buffer.Bytes())
	}

/**
 * Attempts to decode the given address assuming it is a base58 address.
 */

	func decodeAddress(addr string) (dec decod,err error) {
		a,err:=base58check.Decode(addr)
		if err!=nil{
			return dec,err
		}
		versByte := a[0]
		sl := a[1:]
		var hh []uint8
		for _,t := range sl{
			hh=append(hh,uint8(t))
		}
		dec.hash = hh
		switch versByte {
		case 0:
			dec.tp="P2PKH"
			dec.network="mainnet"
			dec.format="legacy"
		case 5:
			dec.tp="P2SH"
			dec.network="testnet"
			dec.format="legacy"
		case 111:
			dec.tp="P2PKH"
			dec.network="testnet"
			dec.format="legacy"
		default:
			return dec,errors.New("Wrong versByte")
		}
		return dec,nil

	}

/**
* Attempts to decode the given address assuming it is a cashaddr address.
*/
	func decodeCashAddress (address string)(dec decod, err error) {
		if (funk.IndexOf(address,":") != -1) {
				return decodeCashAddressWithPrefix(address)
		}else{
			return decod{},errors.New("Missed prefix")
		}
		/*else {
			 prefixes := []string{"bitcoincash", "bchtest", "bchreg"}
			for  i := 0; i < len(prefixes); i++ {
					var prefix = prefixes[i]
					return decodeCashAddressWithPrefix(prefix + ':' + address)
			}
		}*/
	}


/**
* Attempts to decode the given address assuming it is a cashaddr address with explicit prefix.Ы
*/
	func decodeCashAddressWithPrefix (address string)(dec decod, err error) {
		decoded, err := decode(address)
		dec.format="cashaddr"
		dec.tp=decoded.tp
		dec.hash=decoded.hash
		switch (decoded.prefix) {
		case "bitcoincash":
			dec.network="mainnet"
		case "bchtest":
			dec.network="testnet"
		case "bchreg":
			dec.network="regtest"
		default:
			err=errors.New("Wrong prefix")
		}
		return dec,nil
	}













