package db

import "example/web3-profile-api/model"

var UsersTable = []model.User{
	{
		"1",
		"Ridvan Sumset",
		"https://twitter.com/ridvan-sumset",
		"0x1B09a5Bfb014672612e9B4eF2b5592123737fFga",
		450,
		[]string{"diamond"},
		"Ridvan Sumset is arguably the best basketball player ever existed. His score on one-on-one games against MJ 7-1 & against LeBron 9-0. His comment on these games was just \"eZ\".",
	},
	{
		"2",
		"LeBron James",
		"https://twitter.com/lebron-james",
		"0x2B09a5Bfb014672612e9B4eF2b559212373723ty",
		320,
		[]string{"hat", "medal"},
		"LeBron Raymone James Sr. is an American professional basketball player for the Los Angeles Lakers of the National Basketball Association.",
	},
	{
		"3",
		"Michael Jordan",
		"https://twitter.com/michael-jordan",
		"0x1C09a5Bfb014672612e9B4eF2b5592123737u2gt",
		660,
		[]string{"diamond", "hat", "medal"},
		"Michael Jeffrey Jordan, also known by his initials MJ, is an American businessman and former professional basketball player. His biography on the official National Basketball Association website states: \"By acclamation, Michael Jordan is the greatest basketball player of all time.\"",
	},
}
