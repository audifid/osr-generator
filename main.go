package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Char struct {
	name  string
	class string
	align string
	lvl   int
	xp    int
	Params
	Saves
	Combat
	Skills
	Equipment
}

type Params struct {
	Strength
	Dexterity
	Endurance
	Intelligence
	Wisdom
	Charisma
	prime int
}

type Saves struct {
	d int
	w int
	p int
	b int
	s int
}

type Combat struct {
	ac    int
	thac0 int
	hp    int
	hd    int
	dmg   int
}

type Equipment []string

type Skills struct {
	armor  string
	weapon string
	Thief
	Cleric
	Mage
}

type Thief struct {
	useScrolls bool
	backStab   bool
	readLangs  bool
	vl         int
	pl         int
	pr         int
	sk         int
	tp         int
	vz         int
	kk         int
}

type Cleric struct {
	undead int
	dm1    int
	dm2    int
	dm3    int
	dm4    int
	dm5    int
}

type Mage struct {
	mm1 int
	mm2 int
	mm3 int
	mm4 int
	mm5 int
	mm6 int
}

type Strength struct {
	str       int
	melee     int
	openDoors int
}

type Dexterity struct {
	dex        int
	acmod      int
	missile    int
	initiative int
}

type Endurance struct {
	end   int
	hpmod int
}

type Intelligence struct {
	int             int
	spokenLanguages float64
	literacy        string
}

type Wisdom struct {
	wis        int
	magicSaves int
}

type Charisma struct {
	cha          int
	reaction     int
	maxRetainers int
	loyalty      int
}

func (c *Char) CreateChar() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите имя персонажа: ")
	input, _ := reader.ReadString('\n')
	c.name = strings.TrimSpace(input)

	fmt.Println("Выберите класс персонажа: (1-4)")
	fmt.Println("1. Воин")
	fmt.Println("2. Священник")
	fmt.Println("3. Вор")
	fmt.Println("4. Маг")
	fmt.Print("Ваш выбор: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "1" {
		c.class = "Воин"
	} else if input == "2" {
		c.class = "Священник"
	} else if input == "3" {
		c.class = "Вор"
	} else if input == "4" {
		c.class = "Маг"
	} else {
		fmt.Println("Неправильный выбор класса")
	}

	fmt.Println("Выберите мировоззрение: (1-3)")
	fmt.Println("1. Законное")
	fmt.Println("2. Нейтральное")
	fmt.Println("3. Хаотичное")
	fmt.Print("Ваш выбор: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "1" {
		c.align = "Законное"
	} else if input == "2" {
		c.align = "Нейтральное"
	} else if input == "3" {
		c.align = "Хаотичное"
	} else {
		fmt.Println("Неправильный выбор мировоззрения")
	}

	for true {
		str := rand.Intn(16) + 3
		dex := rand.Intn(16) + 3
		end := rand.Intn(16) + 3
		int := rand.Intn(16) + 3
		wis := rand.Intn(16) + 3
		cha := rand.Intn(16) + 3

		fmt.Println("Характеристики:")
		fmt.Printf("Сила: %d\n", str)
		fmt.Printf("Ловкость: %d\n", dex)
		fmt.Printf("Телосложение: %d\n", end)
		fmt.Printf("Интеллект: %d\n", int)
		fmt.Printf("Мудрость: %d\n", wis)
		fmt.Printf("Харизма: %d\n", cha)

		fmt.Print("Принять? (y/n) ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "y" {
			c.str = str
			c.dex = dex
			c.end = end
			c.int = int
			c.wis = wis
			c.cha = cha
			break
		}
	}

	if c.class == "Воин" {
		c.d = 12
		c.w = 13
		c.p = 14
		c.b = 15
		c.s = 16
	} else if c.class == "Священник" {
		c.d = 11
		c.w = 12
		c.p = 14
		c.b = 16
		c.s = 15
	} else if c.class == "Вор" {
		c.d = 13
		c.w = 14
		c.p = 13
		c.b = 16
		c.s = 15
	} else if c.class == "Маг" {
		c.d = 13
		c.w = 14
		c.p = 13
		c.b = 16
		c.s = 15
	}

	// Prime
	var main int

	if c.class == "Воин" {
		main = c.str
	} else if c.class == "Священник" {
		main = c.wis
	} else if c.class == "Вор" {
		main = c.dex
	} else if c.class == "Маг" {
		main = c.int
	}

	if main >= 3 && main <= 5 {
		c.prime = -20
	} else if main >= 6 && main <= 8 {
		c.prime = -10
	} else if main >= 9 && main <= 12 {
		c.prime = 0
	} else if main >= 13 && main <= 15 {
		c.prime = 5
	} else if main >= 16 && main <= 18 {
		c.prime = 10
	}

	c.xp = 0
	c.lvl = 1

	fmt.Println("")
	fmt.Printf("Имя: %s\n", c.name)
	fmt.Printf("Класс: %s\n", c.class)
	fmt.Printf("Мировоззрение: %s\n", c.align)
	fmt.Printf("Уровень: %d, Опыт: %d\n", c.lvl, c.xp)

	fmt.Println("")
	fmt.Printf("СИЛА:%d ЛОВК:%d ТЕЛО:%d ИНТ:%d МДР:%d ХАР:%d \n", c.str, c.dex, c.end, c.int, c.wis, c.cha)
	fmt.Printf("СМЕРТЬ:%d ЖЕЗЛЫ:%d ПАРАЛИЧ:%d ДЫХАНИЕ:%d ЗАКЛИНАНИЯ:%d \n", c.d, c.w, c.p, c.b, c.s)
	fmt.Printf("Бонус к опыту: %d процентов \n", c.prime)
}

func (c *Char) MakeChar() {

	//Strength
	if c.str == 3 {
		c.melee = -3
		c.openDoors = 1
	} else if c.str >= 4 && c.str <= 5 {
		c.melee = -2
		c.openDoors = 1
	} else if c.str >= 6 && c.str <= 8 {
		c.melee = -1
		c.openDoors = 1
	} else if c.str >= 9 && c.str <= 12 {
		c.melee = 0
		c.openDoors = 2
	} else if c.str >= 13 && c.str <= 15 {
		c.melee = 1
		c.openDoors = 3
	} else if c.str >= 16 && c.str <= 17 {
		c.melee = 2
		c.openDoors = 4
	} else if c.str == 18 {
		c.melee = 3
		c.openDoors = 5
	}

	// Dexterity
	if c.dex == 3 {
		c.acmod = -3
		c.missile = -3
		c.initiative = -2
	} else if c.dex >= 4 && c.dex <= 5 {
		c.acmod = -2
		c.missile = -2
		c.initiative = -1
	} else if c.dex >= 6 && c.dex <= 8 {
		c.acmod = -1
		c.missile = -1
		c.initiative = -1
	} else if c.dex >= 9 && c.dex <= 12 {
		c.acmod = 0
		c.missile = 0
		c.initiative = 0
	} else if c.dex >= 13 && c.dex <= 15 {
		c.acmod = 1
		c.missile = 1
		c.initiative = 1
	} else if c.dex >= 16 && c.dex <= 17 {
		c.acmod = 2
		c.missile = 2
		c.initiative = 1
	} else if c.dex == 18 {
		c.acmod = 3
		c.missile = 3
		c.initiative = 2
	}

	// Endurance
	if c.end == 3 {
		c.hpmod = -3
	} else if c.end >= 4 && c.end <= 5 {
		c.hpmod = -2
	} else if c.end >= 6 && c.end <= 8 {
		c.hpmod = -1
	} else if c.end >= 9 && c.end <= 12 {
		c.hpmod = 0
	} else if c.end >= 13 && c.end <= 15 {
		c.hpmod = 1
	} else if c.end >= 16 && c.end <= 17 {
		c.hpmod = 2
	} else if c.end == 18 {
		c.hpmod = 3
	}

	// Intelligence
	if c.int == 3 {
		c.spokenLanguages = 0.5
		c.literacy = "illiterate"
	} else if c.int >= 4 && c.int <= 5 {
		c.spokenLanguages = 1
		c.literacy = "illiterate"
	} else if c.int >= 6 && c.int <= 8 {
		c.spokenLanguages = 1
		c.literacy = "basic"
	} else if c.int >= 9 && c.int <= 12 {
		c.spokenLanguages = 1
		c.literacy = "literate"
	} else if c.int >= 13 && c.int <= 15 {
		c.spokenLanguages = 2
		c.literacy = "literate"
	} else if c.int >= 16 && c.int <= 17 {
		c.spokenLanguages = 3
		c.literacy = "literate"
	} else if c.int == 18 {
		c.spokenLanguages = 4
		c.literacy = "literate"
	}

	// Wisdom
	if c.wis == 3 {
		c.magicSaves = -3
	} else if c.wis >= 4 && c.wis <= 5 {
		c.magicSaves = -2
	} else if c.wis >= 6 && c.wis <= 8 {
		c.magicSaves = -1
	} else if c.wis >= 9 && c.wis <= 12 {
		c.magicSaves = 0
	} else if c.wis >= 13 && c.wis <= 15 {
		c.magicSaves = 1
	} else if c.wis >= 16 && c.wis <= 17 {
		c.magicSaves = 2
	} else if c.wis == 18 {
		c.magicSaves = 3
	}

	// Charisma
	if c.cha == 3 {
		c.reaction = -2
		c.maxRetainers = 1
		c.loyalty = 4
	} else if c.cha >= 4 && c.cha <= 5 {
		c.reaction = -1
		c.maxRetainers = 2
		c.loyalty = 5
	} else if c.cha >= 6 && c.cha <= 8 {
		c.reaction = -1
		c.maxRetainers = 3
		c.loyalty = 6
	} else if c.cha >= 9 && c.cha <= 12 {
		c.reaction = 0
		c.maxRetainers = 4
		c.loyalty = 7
	} else if c.cha >= 13 && c.cha <= 15 {
		c.reaction = 1
		c.maxRetainers = 5
		c.loyalty = 8
	} else if c.cha >= 16 && c.cha <= 17 {
		c.reaction = 1
		c.maxRetainers = 6
		c.loyalty = 9
	} else if c.wis == 18 {
		c.reaction = 2
		c.maxRetainers = 7
		c.loyalty = 10
	}

	c.ac = 9 - c.acmod
	if c.ac > 10 {
		c.ac = 10
	}
	c.thac0 = 19 - c.melee
	if c.thac0 > 20 {
		c.thac0 = 20
	}
	c.hd = 1
	for true {
		if c.class == "Воин" {
			c.hp = rand.Intn(8) + 1 + c.hpmod
		} else if c.class == "Священник" {
			c.hp = rand.Intn(6) + 1 + c.hpmod
		} else if c.class == "Вор" {
			c.hp = rand.Intn(4) + 1 + c.hpmod
		} else if c.class == "Маг" {
			c.hp = rand.Intn(4) + 1 + c.hpmod
		}
		if c.hp > 1 {
			break
		}
	}

	// Fighter
	if c.class == "Воин" {
		c.armor = "all"
		c.weapon = "all"
	}
	// Thief
	if c.class == "Вор" {
		c.armor = "max - leather, no shields"
		c.weapon = "all"
		c.useScrolls = false
		c.readLangs = false
		c.backStab = true
		c.vl = 87
		c.pl = 10
		c.pr = 2
		c.sk = 10
		c.tp = 20
		c.vz = 15
		c.kk = 20
	}
	// Cleric
	if c.class == "Священник" {
		c.armor = "all"
		c.weapon = "blunt"
		c.undead = 1
	}
	// Mage
	if c.class == "Маг" {
		c.armor = "none"
		c.weapon = "dagger & staff"
		c.mm1 = 1
	}

}

func (c *Char) SaveChar() error {
	filename := fmt.Sprintf("%s.char", c.name)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	slice := make([]string, 0)
	slice = append(slice, c.name, c.class, c.align)
	slice = append(slice, strconv.Itoa(c.xp), strconv.Itoa(c.lvl), strconv.Itoa(c.prime))
	slice = append(slice, strconv.Itoa(c.d), strconv.Itoa(c.w), strconv.Itoa(c.p), strconv.Itoa(c.b), strconv.Itoa(c.s))
	slice = append(slice, strconv.Itoa(c.str), strconv.Itoa(c.dex), strconv.Itoa(c.end), strconv.Itoa(c.int), strconv.Itoa(c.wis), strconv.Itoa(c.cha))
	for _, v := range slice {
		str := fmt.Sprintf("%s\n", v)
		_, err = file.WriteString(str)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Char) LoadChar() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя персонажа для загрузки: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	input = strings.TrimSpace(input)
	filename := fmt.Sprintf("%s.char", input)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	slice := make([]string, 0)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	c.name = slice[0]
	c.class = slice[1]
	c.align = slice[2]
	c.xp, _ = strconv.Atoi(slice[3])
	c.lvl, _ = strconv.Atoi(slice[4])
	c.prime, _ = strconv.Atoi(slice[5])

	c.d, _ = strconv.Atoi(slice[6])
	c.w, _ = strconv.Atoi(slice[7])
	c.p, _ = strconv.Atoi(slice[8])
	c.b, _ = strconv.Atoi(slice[9])
	c.s, _ = strconv.Atoi(slice[10])

	c.str, _ = strconv.Atoi(slice[11])
	c.dex, _ = strconv.Atoi(slice[12])
	c.end, _ = strconv.Atoi(slice[13])
	c.int, _ = strconv.Atoi(slice[14])
	c.wis, _ = strconv.Atoi(slice[15])
	c.cha, _ = strconv.Atoi(slice[16])

	return nil
}

func (c *Char) ViewChar() {
	fmt.Println("")
	fmt.Printf("Имя: %s\n", c.name)
	fmt.Printf("Класс: %s\n", c.class)
	fmt.Printf("Мировоззрение: %s\n", c.align)
	fmt.Printf("Уровень: %d, Опыт: %d\n", c.lvl, c.xp)

	fmt.Println("")
	fmt.Printf("СИЛА:%d ЛОВК:%d ТЕЛО:%d ИНТ:%d МДР:%d ХАР:%d \n", c.str, c.dex, c.end, c.int, c.wis, c.cha)
	fmt.Printf("СМЕРТЬ:%d ЖЕЗЛЫ:%d ПАРАЛИЧ:%d ДЫХАНИЕ:%d ЗАКЛИНАНИЯ:%d \n", c.d, c.w, c.p, c.b, c.s)
	fmt.Printf("Бонус к опыту: %d процентов \n", c.prime)

	fmt.Println("")
	fmt.Printf("Бонус к атаке и урону в ближнем бою: %d\n", c.melee)
	fmt.Printf("Выбивание дверей %d\n", c.openDoors)
	fmt.Printf("Бонус к AC: %d\n", c.acmod)
	fmt.Printf("Бонус к атаке в дальнем бою: %d\n", c.missile)
	fmt.Printf("Инициатива: %d\n", c.initiative)
	fmt.Printf("Бонус к ХП: %d\n", c.hpmod)
	fmt.Printf("Языки: %.f\n", c.spokenLanguages)
	fmt.Printf("Образованность: %s\n", c.literacy)
	fmt.Printf("Спасброски против магии: %d\n", c.magicSaves)
	fmt.Printf("Реакция: %d\n", c.reaction)
	fmt.Printf("Максимум последователей: %d\n", c.maxRetainers)
	fmt.Printf("Лояльность: %d\n", c.loyalty)

	fmt.Println("")
	fmt.Printf("AC: %d\n", c.ac)
	fmt.Printf("THAC0: %d\n", c.thac0)
	fmt.Printf("HP: %d\n", c.hp)
	fmt.Printf("DMG: %d\n", c.dmg)

	fmt.Println("")
	fmt.Printf("Доспехи: %s\n", c.armor)
	fmt.Printf("Оружие: %s\n", c.weapon)

	if c.class != "Воин" {
		fmt.Println("")
		fmt.Println("Способности:")
	}
	if c.class == "Вор" {
		if c.backStab == true {
			fmt.Printf("Удар в спину\n")
		}
		if c.useScrolls == true {
			fmt.Printf("Использование свитков\n")
		}
		if c.readLangs == true {
			fmt.Printf("Чтение языков\n")
		}
		fmt.Printf("Лазание по вертикальным поверхностям: %d\n", c.vl)
		fmt.Printf("Обезвреживание ловушек на сокровищах: %d\n", c.pl)
		fmt.Printf("Прислушивание: %d\n", c.pr)
		fmt.Printf("Скрыться в тенях: %d\n", c.sk)
		fmt.Printf("Тихое передвижение: %d\n", c.tp)
		fmt.Printf("Вскрытие замков: %d\n", c.vz)
		fmt.Printf("Карманная кража: %d\n", c.kk)
	} else if c.class == "Священник" {
		fmt.Printf("Изгнание нежити: %d\n", c.undead)
		fmt.Printf("Заклинания 1-го уровня: %d\n", c.dm1)
	} else if c.class == "Маг" {
		fmt.Printf("Заклинания 1-го уровня: %d\n", c.mm1)
	}

	fmt.Println("")
	fmt.Printf("Снаряжение: %v\n", c.Equipment)
}

func main() {
	c := &Char{}
	fmt.Println("Генератор 1lvl персонажей OSR D&D")
	fmt.Println("Выберите опцию: ")
	fmt.Println("1. Создать персонажа / 2. Загрузить персонажа")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ваш выбор: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	if input == "1" {
		c.CreateChar()
		fmt.Println("Сохранить? (y/n): ")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		if input == "y" {
			c.SaveChar()
			c.MakeChar()
			c.ViewChar()
		}
	} else if input == "2" {
		c.LoadChar()
		c.MakeChar()
		c.ViewChar()
	}
}

// Переместить генерацию хп в создание а не makechar
// Убрать hd
// Добавить систему повышения уровня
