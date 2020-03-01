import CryptoJS from '../../../crypto-js'

const keyDefault = CryptoJS.enc.Utf8.parse("1234123412ABCDEF");  //十六位十六进制数作为密钥
const ivDefault = CryptoJS.enc.Utf8.parse('ABCDEF1234123412');   //十六位十六进制数作为密钥偏移量

class crypt {
    constructor(key, iv) {
        this.key = key ? CryptoJS.enc.Utf8.parse(key) : keyDefault
        this.iv = iv ? CryptoJS.enc.Utf8.parse(iv) : ivDefault
    }
    getKey(ha) {
        return [this.key, ha]
    }
    //解密方法
    Decrypt(word) {
        let encryptedHexStr = CryptoJS.enc.Hex.parse(word);
        let srcs = CryptoJS.enc.Base64.stringify(encryptedHexStr);
        let decrypt = CryptoJS.AES.decrypt(srcs, this.key, { iv: this.iv, mode: CryptoJS.mode.CBC, padding: CryptoJS.pad.Pkcs7 });
        let decryptedStr = decrypt.toString(CryptoJS.enc.Utf8);
        return decryptedStr.toString();
    }

    //加密方法
    Encrypt(word) {
        let srcs = CryptoJS.enc.Utf8.parse(word);
        let encrypted = CryptoJS.AES.encrypt(srcs, this.key, { iv: this.iv, mode: CryptoJS.mode.CBC, padding: CryptoJS.pad.Pkcs7 });
        return encrypted.ciphertext.toString().toUpperCase();
    }
}

export default crypt
