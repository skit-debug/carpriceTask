CREATE DATABASE BOOK_CATALOG;
USE BOOK_CATALOG;

CREATE TABLE authors(
author_id INT auto_increment,
first_name CHAR(30) NOT NULL,
last_name CHAR(30) NOT NULL,
born DATETIME NOT NULL,
died DATETIME,
PRIMARY KEY(author_id)
);

CREATE TABLE books(
book_id INT auto_increment,
title CHAR(100) NOT NULL,
release_year INT,
abstract TEXT(1000),
author_id INT,
PRIMARY KEY(book_id),
FOREIGN KEY(author_id) REFERENCES authors(author_id)
);
 
INSERT INTO authors
VALUES
(0, "Joanne", "Rowling", "1965-07-31", NULL),
(0, "Vladimir", "Nabokov", "1899-04-10", "1977-07-02"),
(0, "Viktor", "Pelevin", "1962-11-22", NULL),
(0, "Joseph", "Brodsky", "1940-05-24", "1996-01-28");

INSERT INTO books
VALUES
(0, "Harry Potter and the Philosopher's Stone", 1997, "Harry Potter lives with his abusive aunt and uncle, Vernon and Petunia Dursley and their bullying son, Dudley. On Harry's eleventh birthday, a half-giant named Rubeus Hagrid personally delivers an acceptance letter to Hogwarts School of Witchcraft and Wizardry. Harry's parents, James and Lily Potter, were wizards. When Harry was one year old, an evil and powerful dark wizard, Lord Voldemort, murdered his parents. Harry survived Voldemort's killing curse that rebounded off his forehead and seemingly destroyed the Dark Lord. A lightning bolt-shaped scar was left on his forehead. Unknown to Harry, he is famous in the wizarding world.", 1),
(0, "Harry Potter and the Chamber of Secrets", 1998, "While spending the summer at the Dursleys, twelve-year-old Harry is visited by a house-elf named Dobby. He warns that Harry is in danger and must not return to Hogwarts. Harry refuses, so Dobby magically ruins Aunt Petunia and Uncle Vernon's dinner party. A furious Uncle Vernon locks Harry into his room in retaliation. The Ministry of Magic immediately sends a notice accusing Harry of performing underage magic and threatening dismissal from Hogwarts.", 1),
(0, "Harry Potter and the Prisoner of Azkaban", 1999, "Thirteen-year-old Harry Potter spends another unhappy summer at the Dursleys. After Aunt Marge insults Harry and his deceased parents, an angry Harry accidentally inflates her. Fearing expulsion from Hogwarts, he runs away. On a dark street, a large black dog watches Harry. Startled, Harry stumbles backward, causing his wand to emit sparks. The Knight Bus, a vehicle that rescues stranded wizards, suddenly arrives. Harry goes to the Leaky Cauldron in Diagon Alley where Cornelius Fudge, the Minister for Magic, is waiting. Harry is not expelled but is asked to remain in Diagon Alley until school starts. While there, Harry reunites with best friends Ron Weasley and Hermione Granger. Mr Weasley warns Harry about the wizard Sirius Black, a convicted murderer who escaped Azkaban prison and is believed to be hunting down Harry.", 1),
(0, "Harry Potter and the Goblet of Fire", 2000, "Throughout the three previous novels in the Harry Potter series, the main character, Harry Potter, has struggled with the difficulties of growing up and the added challenge of being a famed wizard. When Harry was a baby, Lord Voldemort, the most powerful dark wizard in history, killed Harry's parents but was mysteriously defeated after unsuccessfully trying to kill Harry, though his attempt left a lightning-shaped scar on Harry's forehead. This results in Harry's immediate fame and his being placed in the care of his abusive Muggle (non-magical) aunt and uncle, Petunia and Vernon Dursley, who have a son named Dudley.", 1),
(0, "Harry Potter and the Order of the Phoenix", 2003, "During the summer, Harry Potter and his cousin Dudley are attacked by Dementors. Forced to magically fend them off, Harry is expelled from Hogwarts, but his expulsion is postponed pending a hearing at the Ministry of Magic. A group of wizards belonging to the Order of the Phoenix whisk Harry off to Number 12, Grimmauld Place, Sirius Black's childhood home.", 1),
(0, "Harry Potter and the Half-Blood Prince", 2005, "Severus Snape, a member of the Order of the Phoenix, meets with Narcissa Malfoy, Draco's mother, and her sister Bellatrix Lestrange, Lord Voldemort's supporter. Narcissa expresses concern that her son may not survive a mission that Voldemort has given him. Snape makes an Unbreakable Vow with Narcissa, swearing to assist Draco.", 1),
(0, "Harry Potter and the Deathly Hallows", 2007, "Throughout the six previous novels in the series, the main character Harry Potter has struggled with the difficulties of adolescence along with being famous as the only person ever to survive the Killing Curse. The curse was cast by Tom Riddle, better known as Lord Voldemort, a powerful evil wizard who murdered Harry's parents and attempted to kill Harry as a baby, due to a prophecy which claimed Harry would be able to stop him. As an orphan, Harry was placed in the care of his Muggle (non-magical) relatives Petunia Dursley and Vernon Dursley, with their son Dudley Dursley.", 1),
(0, "Lolita", 1955, "The novel is prefaced by a fictitious foreword by John Ray Jr., an editor of psychology books. Ray states that he is presenting a memoir written by a man using the pseudonym Humbert Humbert, who had recently died of heart disease while awaiting a murder trial in jail. The memoir, which addresses the audience as his jury, begins with Humbert's birth in Paris in 1910. He spends his childhood on the French Riviera, where he falls in love with his friend Annabel Leigh. This youthful and physically unfulfilled love is interrupted by Annabel's premature death from typhus, which causes Humbert to become sexually obsessed with a specific type of girl, aged 9 to 14, whom he refers to as nymphets.", 2),
(0, "Chapayev and Void", 1996, "The book is set in two different times — after the October Revolution and in modern Russia. In the post-revolutionary period, Pyotr Pustota is a poet who has fled from Saint Petersburg to Moscow and who takes up the identity of a Soviet political commissar and meets a strange man named Vasily Chapayev who is some sort of an army commander. He spends his days drinking samogon, taking drugs and talking about the meaning of life with Chapayev.", 3),
(0, "Generation  P", 1999, "The novel is set in Moscow in the Yeltsin years, the early 1990s, a time of rampant chaos and corruption. Its protagonist, Babylen Tatarsky, graduate student and poet, has been tossed onto the streets after the fall of the Soviet Union where he soon learns his true calling: developing Russian versions of western advertisements. But the more he succeeds as a copywriter, the more he searches for meaning in a culture now defined by material possessions and self-indulgence. He attempts to discover the forces that determine individual desires and shape collective belief in this post-Soviet world. In this quest, Tatarsky sees coincidences that suggest patterns that in turn suggest a hidden meaning behind the chaos of life. He first senses this hidden purpose when reading about Mesopotamian religious practices. Tatarsky's quest is enhanced by the consumption of hallucinogenic mushrooms, cocaine, and vodka. His quest is further aided by another form of spirits: through a ouija board, Che Guevara writes a treatise on identity, consumerism, and television. Eventually, Tatarsky begins to learn some truths—for instance, that all of politics and the “real” events broadcast on television are digital creations. But he can never quite discover the ultimate force behind these fabrications. When at last he reaches the top of the corporate pyramid, Tatarsky learns that the members of his firm are servants of the goddess Ishtar, whose corporeal form consists of the totality of advertising images. The firm's chief duty is to make sure that Ishtar's enemy, the dog Phukkup, does not awaken, bringing with it chaos and destruction. After a ritual sacrifice, Tatarsky becomes the goddesses’ new regent and, in the form of a 3-D double, her bridegroom. In the novel's last chapter, Tatarsky's electronic double becomes a ubiquitous presence on Russian TV. Tatarsky, who had tried to look past the false images presented on TV to see a true unmediated reality, has himself been transformed into an illusion.", 3),
(0, "Empire V", 2006, "A young man named Roman Aleksandrovich Shtorkin becomes a vampire. This happens when Roman accidentally meets another vampire, Brama, who decided to kill himself after a vampire duel. But before he does, he is obliged to give the other man his tongue - the special essence that makes a person a vampire. With the help of the tongue, a vampire can read the mind of a human or another vampire by tasting their blood. As vampires say, tasting.", 3),
(0, "Selected Poems", 1973, "", 4);

SHOW databases;
SHOW tables;
SELECT *
FROM authors;
SELECT *
FROM books;
DROP TABLE books;
DROP TABLE authors;