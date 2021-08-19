package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseInt(t *testing.T) {
	actual := reverseInt(123)
	if actual != 321 {
		t.Errorf("expected: 321, actual: %d", actual)
	}

	actual = reverseInt(0)
	if actual != 0 {
		t.Errorf("expected: 0, actual: %d", actual)
	}

	actual = reverseInt(444)
	if actual != 444 {
		t.Errorf("expected: 444, actual: %d", actual)
	}

	actual = reverseInt(9223372036854775199)
	if actual != 0 {
		t.Errorf("expected: 0, actual: %d", actual)
	}
}

// Тест описывает различие методов t.Fail и t.Fatal
// когда t.Fail просто помечает тест как не прошедгий но продолжает дальнейшие проверки
// то t.Fatal завершает выполнение тестовой функции. Для того чтобы проверить это замените
// в каждой проверке 321 на 3211 и запустите тест. Сообщение, которое вы должны увидеть в консоли - "case 2"
func TestInvalidReverseInt(t *testing.T) {
	actual := reverseInt(123)
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
	req := require.New(t)
	req.Equal(321, reverseInt(123))
	req.Equal(0, reverseInt(0))
	req.Equal(444, reverseInt(444))
	req.Equal(0, reverseInt(9223372036854775199))
}
