package i18n

import (
	"os"
	"testing"
)

func TestGetLanguage(t *testing.T) {
	// 保存原始环境变量
	originalLang := os.Getenv("LANG")
	originalLcAll := os.Getenv("LC_ALL")

	// 清理环境变量
	os.Unsetenv("LANG")
	os.Unsetenv("LC_ALL")

	// 默认情况下应返回英文
	defaultLang := GetLanguage()
	if defaultLang != EN {
		t.Errorf("Expected default language to be EN, got %s", defaultLang)
	}

	// 设置语言并测试
	SetLanguage(ZH)
	if GetLanguage() != ZH {
		t.Errorf("Expected language to be ZH after SetLanguage(ZH)")
	}

	SetLanguage(EN)
	if GetLanguage() != EN {
		t.Errorf("Expected language to be EN after SetLanguage(EN)")
	}

	// 测试系统语言获取
	os.Setenv("LANG", "zh_CN.UTF-8")
	if GetSystemLanguage() != ZH {
		t.Errorf("Expected system language to be ZH when LANG=zh_CN.UTF-8")
	}

	os.Setenv("LANG", "en_US.UTF-8")
	if GetSystemLanguage() != EN {
		t.Errorf("Expected system language to be EN when LANG=en_US.UTF-8")
	}

	// 恢复原始环境变量
	os.Setenv("LANG", originalLang)
	os.Setenv("LC_ALL", originalLcAll)

	// 确保恢复默认设置
	SetLanguage(EN) // 重置为默认英文
}

func TestT(t *testing.T) {
	// 测试中文翻译
	zhTranslation := T(ZH, "https_server_title")
	if zhTranslation != "HTTPS 文件服务器 - 让文件分享变得简单而安全" {
		t.Errorf("Expected Chinese translation, got: %s", zhTranslation)
	}

	// 测试英文翻译
	enTranslation := T(EN, "https_server_title")
	if enTranslation != "HTTPS File Server - Making file sharing simple and secure" {
		t.Errorf("Expected English translation, got: %s", enTranslation)
	}

	// 测试未定义的键
	undefinedKey := T(ZH, "undefined_key")
	if undefinedKey != "undefined_key" {
		t.Errorf("Expected undefined key to return itself, got: %s", undefinedKey)
	}
}
