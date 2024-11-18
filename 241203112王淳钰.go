package main

import (
    "fmt"
    "math/rand"
    "time"
)

// 玩家结构体
type Player struct {
    Name    string
    Health  int
    Attack  int
}

// 怪兽结构体
type Monster struct {
    Name    string
    Health  int
    Attack  int
}

// 玩家攻击怪兽的函数
func (p *Player) AttackMonster(m *Monster) {
    damage := rand.Intn(p.Attack) + 1
    m.Health -= damage
    fmt.Printf("%s对%s造成了%d点伤害，%s剩余生命值：%d\n", p.Name, m.Name, damage, m.Name, m.Health)
}

// 怪兽攻击玩家的函数
func (m *Monster) AttackPlayer(p *Player) {
    damage := rand.Intn(m.Attack) + 1
    p.Health -= damage
    fmt.Printf("%s对%s造成了%d点伤害，%s剩余生命值：%d\n", m.Name, p.Name, damage, p.Name, p.Health)
}

func main() {
    rand.Seed(time.Now().UnixNano())

    // 创建玩家
    player := Player{
        Name:    "勇士",
        Health:  100,
        Attack:  20,
    }

    // 创建怪兽
    monster := Monster{
        Name:    "邪恶巨兽",
        Health:  80,
        Attack:  15,
    }

    fmt.Printf("%s出现了，准备战斗吧，%s！\n", monster.Name, player.Name)

    for {
        // 玩家攻击怪兽
        player.AttackMonster(&monster)
        if monster.Health <= 0 {
            fmt.Printf("%s击败了%s，胜利！\n", player.Name, monster.Name)
            break
        }

        // 怪兽攻击玩家
        monster.AttackPlayer(&player)
        if player.Health <= 0 {
            fmt.Printf("%s被%s击败了，失败！\n", player.Name, monster.Name)
            break
        }
    }
}