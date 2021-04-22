package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseInt(t *testing.T) {
	actual := ReverseInt(123)
	if actual != 321 {
		t.Errorf("expected: 321, actual: %d", actual)
	}

	actual = ReverseInt(0)
	if actual != 0 {
		t.Errorf("expected: 0, actual: %d", actual)
	}

	actual = ReverseInt(444)
	if actual != 444 {
		t.Errorf("expected: 444, actual: %d", actual)
	}

	const maxInt32 = 2147483647
	actual = ReverseInt(maxInt32 + 1)
	if actual != 0 {
		t.Errorf("expected: 0, actual: %d", actual)
	}
}

// Тест описывает различие методов t.Fail и t.Fatal
// когда t.Fail просто помечает тест как не прошедгий но продолжает дальнейшие проверки
// то t.Fatal завершает выполнение тестовой функции. Для того чтобы проверить это замените
// в каждой проверке 321 на 3211 и запустите тест. Сообщение, которое вы должны увидеть в консоли - "case 2"
func TestInvalidReverseInt(t *testing.T) {
	actual := ReverseInt(123)
	if actual != 321 {
		t.Fail()
	}

	if actual != 321 {
		t.Fatal("case 2")
	}

	if actual != 321 {
		t.Fatal("case 3")
	}
}

func TestReverseIntV2(t *testing.T) {
	const maxInt32 = 2147483647
	req := require.New(t)
	req.Equal(321, ReverseInt(123))
	req.Equal(0, ReverseInt(0))
	req.Equal(444, ReverseInt(444))
	req.Equal(0, ReverseInt(maxInt32+1))
}
