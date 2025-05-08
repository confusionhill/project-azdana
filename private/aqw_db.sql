/*
SQLyog Community Edition- MySQL GUI v7.11 
MySQL - 5.1.37 : Database - winquestworlds
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

CREATE DATABASE /*!32312 IF NOT EXISTS*/`winquestworlds` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `winquestworlds`;

/*Table structure for table `equipment` */

DROP TABLE IF EXISTS `equipment`;

CREATE TABLE `equipment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `itemID` int(11) NOT NULL,
  `sLink` varchar(32) NOT NULL,
  `sElmt` varchar(16) NOT NULL DEFAULT 'None',
  `bStaff` tinyint(1) NOT NULL DEFAULT '0',
  `iRng` smallint(8) NOT NULL DEFAULT '10',
  `iDPS` smallint(8) NOT NULL DEFAULT '100',
  `bCoins` tinyint(1) NOT NULL,
  `sES` varchar(16) NOT NULL,
  `sType` varchar(16) NOT NULL,
  `iCost` int(16) NOT NULL,
  `iRty` tinyint(1) NOT NULL,
  `iLvl` smallint(3) NOT NULL,
  `sIcon` varchar(16) NOT NULL,
  `iQty` tinyint(3) NOT NULL,
  `iEnh` smallint(3) NOT NULL,
  `iHrs` smallint(3) NOT NULL,
  `sFile` varchar(128) NOT NULL,
  `iStk` tinyint(1) NOT NULL,
  `sDesc` text NOT NULL,
  `bUpg` tinyint(1) NOT NULL,
  `sName` varchar(32) NOT NULL,
  `bTemp` tinyint(1) NOT NULL DEFAULT '0',
  `sFaction` varchar(16) NOT NULL DEFAULT 'None',
  `iClass` int(11) NOT NULL DEFAULT '0',
  `FactionID` int(11) NOT NULL DEFAULT '1',
  `iReqRep` int(11) NOT NULL DEFAULT '0',
  `iReqCP` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=85 DEFAULT CHARSET=latin1;

/*Data for the table `equipment` */

insert  into `equipment`(`id`,`itemID`,`sLink`,`sElmt`,`bStaff`,`iRng`,`iDPS`,`bCoins`,`sES`,`sType`,`iCost`,`iRty`,`iLvl`,`sIcon`,`iQty`,`iEnh`,`iHrs`,`sFile`,`iStk`,`sDesc`,`bUpg`,`sName`,`bTemp`,`sFaction`,`iClass`,`FactionID`,`iReqRep`,`iReqCP`) values (2,1,'Sword01','None',0,10,100,0,'Weapon','Sword',100,0,1,'iwsword',1,1,50,'items/swords/sword01.swf',1,'Sword',0,'Default Sword',0,'None',0,1,0,0),(3,451,'Wings4','None',0,10,100,0,'ba','Cape',100000,0,1,'iicape',1,1,50,'items/capes/wings4.swf',1,'These wings show your focus and determination to receive the things you desire.',0,'White Feather Wings',0,'None',0,1,0,0),(4,280,'Ears1','None',0,10,100,0,'he','Helm',1000,0,1,'iihelm',1,1,50,'items/helms/ears1.swf',1,'Hah! You have ears on top of your head!',0,'Cat Ears',0,'None',0,1,0,0),(5,774,'Peasant2','None',0,104,100,0,'co','Armor',2,0,1,'iwarmor',1,103,50,'peasant2_skin.swf',0,'',0,'Peasant Rags',0,'None',0,1,0,0),(6,572,'Wyvernpet','None',0,103,100,0,'pe','Pet',10000,0,1,'iipet',1,103,50,'items/pets/wyvernpet.swf',1,'Looks like this Wyvern is friendly.',1,'Wyvern Pet',0,'None',0,1,0,0),(7,17,'Tomix','None',0,10,100,0,'ar','Class',900001,0,5,'iiclass',0,104,0,'Tomix_skin.swf',1,'VIP armour.',1,'Bounty Hunter',0,'None',0,1,0,0),(10,18,'gravityrocks1','None',0,10,100,0,'ba','Cape',9001,0,1,'iicape',0,104,0,'items/capes/gravityrocks1.swf',1,'Defying gravity one rock at a time.',1,'Gravity Rocks',0,'None',0,1,0,0),(12,19,'Ragestar','None',0,10,100,0,'Weapon','Sword',20000,0,5,'iwsword',0,1,0,'items/swords/ragestar.swf',1,'Rage star.',0,'Rage Star',0,'None',0,1,0,0),(13,2,'PhoenixHeart','None',0,60,100,0,'Weapon','Sword',20000,0,1,'iwsword',0,1,0,'items/swords/PhoenixHeart.swf',1,'Phoenix Heart made by Slappy. ',0,'Phoenix Heart',0,'None',0,1,0,0),(14,3,'Potatosack','None',0,10,100,0,'ar','Class',9000,0,1,'iiclass',0,1,0,'potatosack_skin.swf',1,'A potato sack for niggers.',0,'Potato Sack',0,'None',0,1,0,0),(15,5,'Tree','None',0,20,100,0,'Weapon','Sword',10,0,1,'iwsword',0,1,0,'items/swords/Tree.swf',1,'My dad had a job of pulling out trees from the ground cuz he was rlly strong. He would come home from work and beat me up bcuz work was really stresful but i respected him still cuz he was mighty strong. One time he crushed my lungs and a ambulance came to my house to save me but he wanted me to die so he drove his car into an amublance and died. His memory lives on through the tree of might. Tree of might used to have a light saber also but it lost it in great battle ',0,'Tree of Might',0,'None',0,1,0,0),(16,4,'Dagger2','None',0,15,100,0,'Weapon','Sword',0,0,1,'iwsword',0,1,0,'items/daggers/dagger02.swf',1,'Default dagger.',0,'Dagger',0,'None',0,1,0,0),(18,6,'skullstaff','None',0,20,100,0,'Weapon','Staff',0,0,1,'iwsword',0,1,0,'items/swords/skullstaff.swf',1,'An oversized staff, with a skull.Made by SolitaireMaker',0,'Mighty Skull Staff of Darkness',0,'None',0,1,0,0),(19,7,'Arpurit','None',0,17,100,0,'Weapon','Sword',5000,0,1,'iwsword',0,1,0,'items/swords/Arpurit.swf',1,'Another oversized weapon with weird animation. By SolitaireMaker',0,'Purify Blade of Gold ',0,'None',0,1,0,0),(20,8,'J6helm','None',0,10,100,0,'he','Helm',50000,0,1,'iihelm',0,1,0,'items/helms/J6.swf',1,'J6\'s Helm. VIP Only',1,'J6\'s Helm',0,'None',0,1,0,0),(21,20,'KFC','None',0,100,100,0,'Weapon','Sword',10,0,1,'iwsword',0,104,0,'items/swords/kfc.swf',1,'A KFC Bucket full of love, for niggers.',0,'Bucket-O-KFC',0,'None',0,1,0,0),(23,21,'GayBlade','None',0,15,100,0,'Weapon','Sword',0,0,1,'iwsword',0,1,0,'items/swords/vampsgay.swf',1,'Yet another over sized sword by SolitaireMaker.',0,'Lokinko\'s Shadow Blade',0,'None',0,1,0,0),(24,22,'Fallon_FallenSabre.swf','None',0,4,100,0,'Weapon','Sword',50000,0,1,'iwsword',0,1,0,'items/swords/Fallon_FallenSabre.swf',1,'No description available.',1,'Fallen Sabre',0,'None',0,1,0,0),(25,23,'silver_dicer','None',0,2,100,0,'Weapon','Sword',100000,0,5,'iwsword',0,1,0,'items/swords/5025_SilverDicer_Wep.swf',1,'Especially useful in the art of decapitation. By 5025.',0,'The Silver Dicer',0,'None',0,1,0,0),(26,24,'Dead_Ruhe','None',0,13,100,0,'Weapon','Sword',20000,0,1,'iwsword',0,1,0,'items/swords/Maelstrom_Dead_Ruhe_AQW2.swf',1,'A blade made by Maelstrom.',0,'Dead Ruhe',0,'None',0,1,0,0),(27,25,'Terror','None',0,15,100,0,'co','Armor',50000,0,1,'iwarmor',1,103,50,'Despair_TerrorArmor.swf',1,'This gruesomely scary armor was forged 1000 years ago by Master Terror. How it survived this long is unknown. Made by Despair.',0,'Terror Armor',0,'None',0,1,0,0),(28,26,'Valkira','None',0,10,100,0,'co','Armor',25000,0,5,'iwarmor',0,104,0,'Despair_Valkira_Armor.swf',1,'Forged by the Lord of Wind, Ventus, for his use in the Elemental War. Made by Despair.',0,'Valkira',0,'None',0,1,0,0),(29,27,'Paladin','None',0,10,100,0,'co','Armor',40000,0,1,'iwarmor',0,1,0,'paladin_skin.swf',1,'',0,'Paladin',0,'None',0,1,0,0),(30,28,'Rage','None',0,10,100,0,'co','Armor',100000,0,1,'iwarmor',0,103,0,'fallon_Rage.swf',1,'Desc is unavailable. An armour made by Fallon.',0,'Rage Plate',0,'None',0,1,0,0),(31,29,'BlackStone','None',0,6,100,0,'Weapon','Sword',25000,0,0,'iwsword',1,1,0,'items/swords/Srar_BlackStone.swf',1,'Forged from a strange rock, this blade seems to have a strange effect on sneevils. Made by Srar.',0,'Black Stone',0,'None',0,1,0,0),(33,30,'Blackknightpolearm1a','None',0,13,100,0,'Weapon','Polearm',100000,0,0,'iwsword',0,1,0,'items/polearms/Blackknightpolearm1a.swf',1,'',0,'Black Knight Polearm',0,'None',0,1,0,0),(34,31,'axe10','None',0,17,100,0,'Weapon','Axe',20000,0,0,'iwsword',0,1,0,'items/axes/axe10.swf',1,'',0,'Axe 10',0,'None',0,1,0,0),(56,16,'Warrior','None',0,10,100,0,'ar','Class',0,0,1,'iiclass',0,1,0,'warrior_skin.swf',1,'',0,'Warrior Class',0,'None',0,1,0,0),(57,775,'Nateicblade','None',0,10,100,0,'Weapon','Sword',7000,0,5,'iwsword',0,1,0,'items/swords/Nateic.swf',0,'A blade forged with a great magic called Up2Power. Seemingly to glow with power.',0,'Nateic Blade',0,'None',0,1,0,0),(58,776,'Fiberscythe','None',0,10,100,0,'Weapon','Sword',20000,0,6,'iwsword',0,1,0,'items/swords/Blood Fiber Scythe 2.swf',0,'This scythe has a special fiber entwined into the blade.',0,'Fiber Scythe',0,'None',0,1,0,0),(60,777,'Balance','None',0,10,100,0,'co','Armor',30000,0,1,'iwarmor',0,1,0,'Balance2.swf',0,'The balance set is the manifestation of universal harmony. It embodies the power of all nine elements and the ability to rip your opponents to shreds.',0,'Balance',0,'None',0,1,0,0),(63,779,'Fox','None',0,10,100,0,'co','Armor',1337,0,1,'iwarmor',0,1,0,'AgentFox.swf',0,'The solver of conspiracies and lord of awesomeness, this suit really does pwn.',0,'FBI Suit',0,'None',0,1,0,0),(64,780,'','None',0,10,100,0,'Weapon','Sword',10,0,1,'iwsword',0,1,0,'items/swords/SlayerGuitar.swf',0,'This weapon radiates pure heavy metal.',0,'Electric Guitar',0,'None',0,1,0,0),(65,781,'','None',0,10,100,0,'Weapon','Sword',1000,0,1,'iwsword',0,1,0,'items/swords/DarkBuster.swf',0,'The chosen weapon of Cloud Strife in the world of shadows.',0,'Dark Buster Blade',0,'None',0,1,0,0),(66,782,'','None',0,10,100,0,'Weapon','Sword',1000,0,1,'iwsword',0,1,0,'items/swords/Buster1.swf',0,'The chosen weapon of Cloud Strife.',0,'Buster Blade',0,'None',0,1,0,0),(67,783,'','None',0,10,100,0,'Weapon','Sword',1000,0,1,'iwsword',0,1,0,'items/swords/ChainsawGun.swf',0,'The weapon of Doom.',0,'Power Chainsaw',0,'None',0,1,0,0),(68,784,'','None',0,10,100,0,'Weapon','Sword',10000,0,1,'iwsword',0,1,0,'items/swords/Anduril.swf',0,'The weapon of the king.',0,'Anduril',0,'None',0,1,0,0),(69,785,'Glasses','None',0,10,100,0,'he','Helm',1000,0,1,'iihelm',0,1,0,'items/helms/Glasses2.swf',0,'The traditional apparel.',0,'FBI Glasses',0,'None',0,1,0,0),(70,786,'Balance','None',0,10,100,0,'he','Helm',1000,0,0,'iihelm',0,1,0,'items/helms/BalanceHelm.swf',0,'The helm of the balance set.',0,'Balance Helm ',0,'None',0,1,0,0),(71,787,'','None',0,10,100,0,'Weapon','Sword',1000,0,1,'iwsword',0,1,0,'items/swords/AracknightPolearm.swf',0,'The weapon of the great Aracknight.',0,'Aracknight Polearm',0,'None',0,1,0,0),(72,788,'CBlade','None',0,10,100,0,'Weapon','Sword',5000,0,1,'iwsword',0,1,0,'items/swords/CircuitBlade.swf',0,'Weapon containing a Magic Circuit, powerfull in the right hands.',0,'Circuit Blade',0,'None',0,1,0,0),(73,789,'OCBlade','None',0,10,100,0,'Weapon','Sword',25000,0,5,'iwsword',0,1,0,'items/swords/OffCircuitBlade.swf',0,'An advanced weapon containing a Magic Circuit, powerfull in the right hands.',0,'Officer\'s Circuit Blade',0,'None',0,1,0,0),(74,790,'','None',0,10,100,0,'Weapon','Sword',1000,0,1,'iwsword',0,1,0,'items/swords/NightmareGenesis.swf',0,'The blade of the turbulent sea.',0,'Aqua Dementia',0,'None',0,1,0,0),(75,791,'FallonHair','None',0,10,100,0,'he','Helm',50000,0,1,'iihelm',0,1,0,'items/helms/FallonHair.swf',0,'If your name is not Fallon, you look very silly right now.',0,'Fallon\'s Hair',0,'None',0,1,0,0),(76,792,'','None',0,10,100,0,'Weapon','Sword',100,0,1,'iwsword',0,1,0,'items/swords/BlackGun.swf',0,'A black gun...',0,'Gun',0,'None',0,1,0,0),(77,793,'','None',0,10,100,0,'Weapon','Sword',1000,0,1,'iwsword',0,1,0,'items/swords/Perifidis.swf',0,'The twilight of the thunder god.',0,'Perfidis',0,'None',0,1,0,0),(78,794,'','None',0,10,100,0,'Weapon','Sword',100,0,1,'iwsword',0,1,0,'items/swords/BoogBlade3.swf',0,'Altered Version.',0,'Boogs Blade',0,'None',0,1,0,0),(79,795,'','None',0,10,100,0,'Weapon','Sword',100,0,1,'iwsword',0,1,0,'items/swords/SasukeKusanagiChokutoChidoriCustom.swf',0,'It tastes, lightningy...',0,'Kusanagi Chokuto',0,'None',0,1,0,0),(80,796,'','None',0,10,100,0,'Weapon','Sword',0,0,1,'iwsword',0,1,0,'items/swords/BalanceWeapon.swf',0,'Test Item, First Draft.',0,'Blade Of Balance',0,'None',0,1,0,0),(84,797,'Wings6','None',0,10,100,0,'ba','Cape',1000,0,1,'iicape',0,1,0,'items/capes/BalanceWings.swf',0,'A true mechanical marvel.',0,'Wings Of Balance',0,'None',0,1,0,0);

/*Table structure for table `friends` */

DROP TABLE IF EXISTS `friends`;

CREATE TABLE `friends` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userid` int(11) NOT NULL,
  `friendid` varchar(128) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=24 DEFAULT CHARSET=latin1;

/*Data for the table `friends` */

/*Table structure for table `items` */

DROP TABLE IF EXISTS `items`;

CREATE TABLE `items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `itemid` int(11) NOT NULL,
  `userid` int(11) NOT NULL,
  `classXP` int(11) DEFAULT NULL,
  `className` varchar(32) DEFAULT NULL,
  `equipped` tinyint(1) NOT NULL DEFAULT '0',
  `sES` varchar(16) NOT NULL,
  `bBank` tinyint(1) NOT NULL DEFAULT '0',
  `iLvl` int(3) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=446 DEFAULT CHARSET=latin1;

/*Data for the table `items` */

insert  into `items`(`id`,`itemid`,`userid`,`classXP`,`className`,`equipped`,`sES`,`bBank`,`iLvl`) values (1,16,1,9473,'Warrior Class',1,'ar',0,6),(2,1,1,NULL,NULL,0,'Weapon',1,1),(34,572,1,NULL,NULL,1,'pe',0,1),(35,774,1,NULL,NULL,0,'co',1,1),(36,280,1,NULL,NULL,0,'he',1,1),(40,451,1,NULL,NULL,0,'ba',1,1),(104,18,1,NULL,NULL,0,'ba',0,1),(105,19,1,NULL,NULL,0,'Weapon',1,1),(149,5,1,NULL,NULL,0,'Weapon',1,1),(196,17,1,9473,NULL,1,'ar',0,5),(197,25,1,NULL,NULL,0,'co',1,1),(198,26,1,NULL,NULL,0,'co',1,1),(199,24,1,NULL,NULL,0,'Weapon',1,1),(224,27,1,NULL,NULL,0,'co',1,1),(398,785,1,NULL,NULL,0,'he',0,1),(399,786,1,NULL,NULL,1,'he',0,1),(400,779,1,NULL,NULL,0,'co',0,1),(401,775,1,NULL,NULL,0,'Weapon',0,1),(402,789,1,NULL,NULL,0,'Weapon',1,1),(403,792,1,NULL,NULL,0,'Weapon',0,1),(412,777,1,NULL,NULL,1,'co',0,1),(413,796,1,NULL,NULL,1,'Weapon',0,1),(414,783,1,NULL,NULL,0,'Weapon',0,1),(415,780,1,NULL,NULL,0,'Weapon',0,1),(417,781,1,NULL,NULL,0,'Weapon',1,1),(445,797,1,NULL,NULL,1,'ba',0,1);

/*Table structure for table `maps` */

DROP TABLE IF EXISTS `maps`;

CREATE TABLE `maps` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `fileName` varchar(128) NOT NULL,
  `monsternumb` varchar(128) NOT NULL,
  `monsterid` varchar(128) NOT NULL,
  `monsterframe` varchar(512) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

/*Data for the table `maps` */

insert  into `maps`(`id`,`name`,`fileName`,`monsternumb`,`monsterid`,`monsterframe`) values (4,'battleon','/Battleon/town-battleon-Aug14.swf','','',''),(5,'yulgar','town-yulgar-darts.swf','','',''),(6,'newbie','/Intro/town-newbie.swf','23,23,23,25,25,25,25,23,165','23,25,165','Fight,Fight2,Fight2,Fight2,Fight2,Boss,Boss,Fight,BookStore');

/*Table structure for table `monsters` */

DROP TABLE IF EXISTS `monsters`;

CREATE TABLE `monsters` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sRace` varchar(32) NOT NULL DEFAULT 'None',
  `MonID` int(11) NOT NULL,
  `intMPMax` int(5) NOT NULL,
  `intGold` int(5) NOT NULL,
  `intLevel` int(3) NOT NULL,
  `strDrops` varchar(128) NOT NULL,
  `intExp` int(5) NOT NULL,
  `iDPS` int(3) NOT NULL,
  `intHPMax` int(5) NOT NULL,
  `strElement` varchar(32) NOT NULL DEFAULT 'None',
  `intRSC` int(3) NOT NULL,
  `strLinkage` varchar(32) NOT NULL,
  `strMonFileName` varchar(128) NOT NULL,
  `strMonName` varchar(64) NOT NULL,
  `intRep` int(3) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

/*Data for the table `monsters` */

insert  into `monsters`(`id`,`sRace`,`MonID`,`intMPMax`,`intGold`,`intLevel`,`strDrops`,`intExp`,`iDPS`,`intHPMax`,`strElement`,`intRSC`,`strLinkage`,`strMonFileName`,`strMonName`,`intRep`) values (1,'None',23,30,33,1,'98:1:1,285:1:0.15,330:1:.65,425:1:0.3,807:1:1',11,100,550,'None',0,'Slimegreen','Slimegreen.swf','Slime',50),(2,'None',25,30,36,2,'102:1:1,37:1:0.1,808:1:.8',22,100,600,'None',0,'Sneevil1','Sneevil1.swf','Sneevil',50),(3,'Orc',165,30,33,1,'810:1:.95',11,100,1500,'None',0,'OrcWarrior1a','OrcWarrior1a.swf','Dogear',0);

/*Table structure for table `servers` */

DROP TABLE IF EXISTS `servers`;

CREATE TABLE `servers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `ip` varchar(15) NOT NULL,
  `count` smallint(5) NOT NULL DEFAULT '0',
  `max` smallint(5) NOT NULL DEFAULT '255',
  `online` tinyint(1) NOT NULL DEFAULT '1',
  `bchat` tinyint(1) NOT NULL DEFAULT '0',
  `ichat` tinyint(1) NOT NULL DEFAULT '0',
  `upgrade` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `servers` */

insert  into `servers`(`id`,`name`,`ip`,`count`,`max`,`online`,`bchat`,`ichat`,`upgrade`) values (1,'WQW','127.0.0.1',0,255,0,1,2,0);

/*Table structure for table `settings` */

DROP TABLE IF EXISTS `settings`;

CREATE TABLE `settings` (
  `name` varchar(64) NOT NULL,
  `version` decimal(4,0) NOT NULL,
  `message` text NOT NULL,
  `newsFile` varchar(64) NOT NULL,
  `mapFile` varchar(64) NOT NULL,
  `bookFile` varchar(64) NOT NULL,
  `xprate` tinyint(1) NOT NULL,
  `goldrate` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `settings` */

insert  into `settings`(`name`,`version`,`message`,`newsFile`,`mapFile`,`bookFile`,`xprate`,`goldrate`) values ('WinQuest Worlds','77','Welcome to WinQuest Worlds!','news/News-Aug14.swf','news/Map-Aug14.swf','news/Book-July16.swf',3,3);

/*Table structure for table `shops` */

DROP TABLE IF EXISTS `shops`;

CREATE TABLE `shops` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `shopid` int(11) NOT NULL,
  `strName` varchar(32) NOT NULL,
  `items` text NOT NULL,
  `sField` varchar(16) NOT NULL DEFAULT '',
  `bStaff` tinyint(1) NOT NULL DEFAULT '0',
  `bHouse` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

/*Data for the table `shops` */

insert  into `shops`(`id`,`shopid`,`strName`,`items`,`sField`,`bStaff`,`bHouse`) values (1,10,'Upgrade Only','451,280,572,18','',0,0),(4,16,'Yulgar Weapon Shop','19,5,24,775,780,781,782,783,784,787,789,790,792,793,794,795,796','',0,0),(2,41,'Armour Shop','774,25,26,27,777,779,785,786,797','',0,0);

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL,
  `password` varchar(32) NOT NULL,
  `access` smallint(2) NOT NULL DEFAULT '5',
  `upgrade` tinyint(1) NOT NULL DEFAULT '0',
  `age` smallint(3) NOT NULL,
  `upgDate` datetime DEFAULT '0000-00-00 00:00:00',
  `upgDays` smallint(3) NOT NULL DEFAULT '0',
  `emailActive` tinyint(1) NOT NULL DEFAULT '5',
  `email` varchar(64) NOT NULL DEFAULT 'none',
  `moderator` tinyint(1) NOT NULL DEFAULT '0',
  `level` smallint(3) NOT NULL DEFAULT '1',
  `cosColorAccessory` int(8) NOT NULL DEFAULT '0',
  `cosColorBase` int(8) NOT NULL DEFAULT '0',
  `cosColorTrim` int(8) NOT NULL DEFAULT '0',
  `plaColorSkin` int(8) NOT NULL DEFAULT '13088131',
  `plaColorHair` int(8) NOT NULL DEFAULT '7027237',
  `plaColorEyes` int(8) NOT NULL DEFAULT '91294',
  `slotBag` int(3) NOT NULL DEFAULT '20',
  `slotBank` tinyint(3) NOT NULL DEFAULT '0',
  `slotHouse` tinyint(3) NOT NULL DEFAULT '20',
  `STR` smallint(3) NOT NULL DEFAULT '0',
  `DEX` smallint(3) NOT NULL DEFAULT '0',
  `INT` smallint(3) NOT NULL DEFAULT '0',
  `END` smallint(3) NOT NULL DEFAULT '0',
  `WIS` smallint(3) NOT NULL DEFAULT '0',
  `LCK` smallint(3) NOT NULL DEFAULT '0',
  `currentClass` smallint(2) NOT NULL DEFAULT '2',
  `xp` int(11) NOT NULL DEFAULT '0',
  `gold` int(11) NOT NULL DEFAULT '10000',
  `coins` int(11) NOT NULL DEFAULT '1000',
  `lastVisited` varchar(32) NOT NULL DEFAULT '',
  `hairID` smallint(3) NOT NULL DEFAULT '52',
  `gender` varchar(1) NOT NULL DEFAULT 'M',
  `hairName` varchar(16) NOT NULL DEFAULT 'Default',
  `hairFile` varchar(128) NOT NULL DEFAULT 'hair/M/Default.swf',
  `curServer` varchar(16) NOT NULL DEFAULT 'WQW',
  `banned` tinyint(1) NOT NULL DEFAULT '0',
  `vip` tinyint(1) NOT NULL DEFAULT '0',
  `ug` tinyint(1) NOT NULL DEFAULT '0',
  `dob` varchar(32) NOT NULL DEFAULT '0/0/0000',
  `signupip` varchar(16) NOT NULL DEFAULT '''''',
  `loginip` varchar(16) NOT NULL DEFAULT '''''',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=latin1;

/*Data for the table `users` */

insert  into `users`(`id`,`username`,`password`,`access`,`upgrade`,`age`,`upgDate`,`upgDays`,`emailActive`,`email`,`moderator`,`level`,`cosColorAccessory`,`cosColorBase`,`cosColorTrim`,`plaColorSkin`,`plaColorHair`,`plaColorEyes`,`slotBag`,`slotBank`,`slotHouse`,`STR`,`DEX`,`INT`,`END`,`WIS`,`LCK`,`currentClass`,`xp`,`gold`,`coins`,`lastVisited`,`hairID`,`gender`,`hairName`,`hairFile`,`curServer`,`banned`,`vip`,`ug`,`dob`,`signupip`,`loginip`) values (1,'admin','5f4dcc3b5aa765d61d8327deb882cf99',50,0,15,'0000-00-00 00:00:00',0,5,'none',1,44,0,0,0,13088131,7027237,91294,20,40,20,0,0,0,0,0,0,2,2361287,2147406553,62000,'battleon-1',52,'M','Default','hair/M/Default.swf','',0,0,0,'0/0/0000','\'\'','127.0.0.1');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
