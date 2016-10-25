package main 
 
// Gods is an exported type that 
// contains a structure for each 
// smite god 
type Gods struct { 
  Name       string 
  Pantheon   string 
  Attack     string 
  Power      string 
  Class      string 
  Difficulty string 
} 
 
var ( 
  GODS []*Gods 
) 
 
func init() { 
  GODS = []*Gods{ 
    &Gods{ 
      Name:       "Agni", 
      Pantheon:   "Hindu", 
      Attack:     "Ranged", 
      Power:      "Magical", 
      Class:      "Mage", 
      Difficulty: "Hard", 
    }, 
    &Gods{ 
      Name:       "Ah Muzen Cab", 
      Pantheon:   "Mayan", 
      Attack:     "Ranged", 
      Power:      "Physical", 
      Class:      "Hunter", 
      Difficulty: "Average", 
    }, 
    &Gods{ 
      Name:       "Ah Puch", 
      Pantheon:   "Mayan", 
      Attack:     "Ranged", 
      Power:      "Magical", 
      Class:      "Mage", 
      Difficulty: "Average", 
    }, 
    &Gods{ 
      Name:       "Amaterasu", 
      Pantheon:   "Japanese", 
      Attack:     "Melee", 
      Power:      "Physical", 
      Class:      "Warrior", 
      Difficulty: "Average", 
    }, 
    &Gods{ 
      Name:       "Anhur", 
      Pantheon:   "Egyptian", 
      Attack:     "Ranged", 
      Power:      "Physical", 
      Class:      "Hunter", 
      Difficulty: "Average", 
    }, 
    &Gods{ 
      Name:       "Anubis", 
      Pantheon:   "Egyptian", 
      Attack:     "Ranged", 
      Power:      "Magical", 
      Class:      "Mage", 
      Difficulty: "Easy", 
    }, 
    &Gods{ 
      Name:       "Ao Kuang", 
      Pantheon:   "Chinese", 
      Attack:     "Melee", 
      Power:      "Magical", 
      Class:      "Mage", 
Initial commit of all work