package fec

import "strings"

type Config struct {
	merchantNo      string
	providerNo      string
	productNo       string
	taskCode        string
	privateKey      string
	systemPublicKey string
}

// MerchantNo 获取商户号
func (c *Config) MerchantNo() string {
	return c.merchantNo
}

// ProviderNo 获取服务商号
func (c *Config) ProviderNo() string {
	return c.providerNo
}

// ProductNo 获取产品编号
func (c *Config) ProductNo() string {
	return c.productNo
}

// TaskCode 获取任务编号
func (c *Config) TaskCode() string {
	return c.taskCode
}

// PrivateKey 获取  RSA 私钥
func (c *Config) PrivateKey() string {
	return c.privateKey
}

// SystemPublicKey 获取系统平台 RSA 公钥
func (c *Config) SystemPublicKey() string {
	return c.systemPublicKey
}

// SetMerchantNo 设置商户号
func (c *Config) SetMerchantNo(merchantNo string) *Config {
	c.merchantNo = strings.TrimSpace(merchantNo)

	return c
}

// SetProviderNo 设置服务商号
func (c *Config) SetProviderNo(providerNo string) *Config {
	c.providerNo = strings.TrimSpace(providerNo)
	return c
}

// SetProductNo 设置产品编号
func (c *Config) SetProductNo(productNo string) *Config {
	c.productNo = strings.TrimSpace(productNo)
	return c
}

// SetTaskCode 设置任务编号
func (c *Config) SetTaskCode(taskCode string) *Config {
	c.taskCode = strings.TrimSpace(taskCode)
	return c
}

// SetPrivateKey 设置 RSA 私钥
func (c *Config) SetPrivateKey(privateKey string) *Config {
	c.privateKey = strings.TrimSpace(privateKey)
	return c
}

// SetSystemPublicKey 设置系统平台 RSA 公钥
// RSA 公钥已默认提供；如需要更换，调用该方法
func (c *Config) SetSystemPublicKey(systemPublicKey string) *Config {
	c.systemPublicKey = strings.TrimSpace(systemPublicKey)
	return c
}

// NewConfig 初始化配置
func NewConfig() *Config {
	c := new(Config)
	c.SetSystemPublicKey("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxIidk18omKtKTMOD5dTdX353yxvwk4J+VwVlQ5/rAJwnFBeXmy3kJG3hJj/K3lChr5wLrIpUJQB/dyQqOxMvklw13Uldwe9nd+ffCZfJ0V7guXVfymnyicd9Dz9leuXLV7H+xmjyrz8pFWdtE1KjN5yMAkFhv8EXPsHNgqMA1yCTysz+z8RCu27BYQC3OFgLLKxGH46gbu6m8kUEbvPJlZ5WtwUGcgv62KOVmg40dVpKhNXPjGaljj2ZVFYJiszBLcVVJ6UzJkCykjdB7BUPxXWuaprStELnGvnKY68fEWME8gRWmMBydUJZ8YNlTf+6gVyHXWw/eGKIC+vUCqKTMwIDAQAB")
	return c
}
