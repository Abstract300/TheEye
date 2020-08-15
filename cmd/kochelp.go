package main

import (
	"bytes"

	"github.com/abstract300/theeye/command"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

func Fortification(ctx *command.Context) error {

	data := [][]string{
		{"0", "Camp", "0", "1"},
		{"1", "Stockage", "50,000", "1.25"},
		{"2", "Rabid Pitbulls", "100,000", "1.63"},
		{"3", "Walled Town", "200,000", "2.07"},
		{"4", "Towers", "400,000", "2.64"},
		{"5", "Battlements", "800,000", "3.37"},
		{"6", "Portcullis", "1,600,000", "4.30"},
		{"7", "Boiling Oil", "3,200,000", "5.48"},
		{"8", "Trenches", "6,400,000", "6.98"},
		{"9", "Moat", "12,800,000", "8.90"},
		{"10", "Drawbridge", "25,600,000", "11.35"},
		{"11", "Fortress", "51,200,000", "14.47"},
		{"12", "Stronghold", "102,400,000", "18.46"},
		{"13", "Palace", "204,800,000", "23.53"},
		{"14", "Keep", "409,600,000", "30.00"},
		{"15", "Citadel", "819,200,000", "38.25"},
		{"16", "Hand of God", "1,638,400,000", "48.77"},
	}

	output := &bytes.Buffer{}
	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"Level", "Fortification", "Cost (in gold)", "Multiplier"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	OutputString := "```\n" + output.String() + "```"
	_, err := ctx.Session.ChannelMessageSend(ctx.ChannelID, OutputString)
	return errors.Wrap(err, "couldn't reply")
}

func Siege(ctx *command.Context) error {
	data := [][]string{
		{"1", "Flaming Arrows", "50,000", "1.30"},
		{"2", "Ballista", "100,000", "1.69"},
		{"3", "Battering Ram", "200,000", "2.20"},
		{"4", "Ladders", "400,000", "2.86"},
		{"5", "Trojan Horse", "800,000", "3.71"},
		{"6", "Catapults", "1,600,000", "4.83"},
		{"7", "War Elephants", "3,200,000", "6.27"},
		{"8", "Siege Towers", "6,400,000", "8.16"},
		{"9", "Trebuchets", "12,800,000", "10.60"},
		{"10", "Black Powder", "25,600,000", "13.79"},
		{"11", "Sappers", "25,600,000", "13.79"},
		{"12", "Dynamite", "102,400,000", "23.30"},
		{"13", "Greek Fire", "204,800,000", "30.29"},
		{"14", "Cannons", "409,600,000", "39.37"},
	}

	output := &bytes.Buffer{}
	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"Level", "Siege", "Cost (in gold)", "Multiplier"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	OutputString := "```\n" + output.String() + "```"
	_, err := ctx.Session.ChannelMessageSend(ctx.ChannelID, OutputString)
	return errors.Wrap(err, "couldn't reply")
}

func Covert(ctx *command.Context) error {
	data := [][]string{
		{"1", "Peeping Tom", "17,500"},
		{"2", "Scout", "35,000"},
		{"3", "Informer", "70,000"},
		{"4", "Operative", "140,000"},
		{"5", "Slueth", "280,000"},
		{"6", "Snoop", "560,000"},
		{"7", "Counterspy", "1,120,000"},
		{"8", "Observer", "2,240,000"},
		{"9", "Mole", "4,480,000"},
		{"10", "Raven", "8,960,000"},
		{"11", "Bagnman", "17,920,000"},
		{"12", "Watcher", "35,840,000"},
		{"13", "Agent", "71,680,000"},
		{"14", "Intelligence", "143,360,000"},
		{"15", "James Bond", "296,720,000"},
		{"16", "All Seeing Eye", "573,440,000"},
		{"17", "Eye of Sight", "1,146,880,000"},
	}

	output := &bytes.Buffer{}
	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"Level", "Covert Skill", "Cost (in gold)"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	OutputString := "```\n" + output.String() + "```"
	_, err := ctx.Session.ChannelMessageSend(ctx.ChannelID, OutputString)
	return errors.Wrap(err, "couldn't reply")
}

func Sentry(ctx *command.Context) error {
	data := [][]string{
		{"1", "Barbed Wire", "17,500"},
		{"2", "Pit & Spikes", "35,000"},
		{"3", "Traps", "70,000"},
		{"4", "Trip Wire", "140,000"},
		{"5", "Hidey Hole", "280,000"},
		{"6", "Guard Dogs", "560,000"},
		{"7", "Lookouts", "1,120,000"},
		{"8", "Sentry Pavilion", "2,240,000"},
		{"9", "Shield Wall", "4,480,000"},
		{"10", "Wolfhound", "8,960,000"},
		{"11", "Tracking Tactics", "17,920,000"},
		{"12", "Traps XXI", "35,840,000"},
		{"13", "Coffee", "71,680,000"},
		{"14", "Mystical Eyepatch", "143,360,000"},
		{"15", "Lavastines Hounds", "296,720,000"},
		{"16", "Though Shalt Not Be Named", "573,440,000"},
		{"17", "Defias Brotherhood", "1,146,880,000"},
	}
	output := &bytes.Buffer{}
	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"Level", "Sentry Skill", "Cost (in gold)"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	OutputString := "```\n" + output.String() + "```"
	_, err := ctx.Session.ChannelMessageSend(ctx.ChannelID, OutputString)
	return errors.Wrap(err, "couldn't reply")
}

func Conscription(ctx *command.Context) error {
	data := [][]string{
		{"0", "25", "0"},
		{"1", "50", "500,000"},
		{"2", "100", "1,000,000"},
		{"3", "200", "3,120,000"},
		{"4", "400", "10,480,000"},
		{"5", "800", "42,000,000"},
		{"6", "1600", "108,468,000"},
		{"7", "3200", "212,680,000"},
		{"8", "6400", "501,340,000"},
		{"9", "12800", "1,260,000,000"},
	}
	output := &bytes.Buffer{}
	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"Consc.Level", "Soldiers", "Cost (in gold)"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	OutputString := "```\n" + output.String() + "```"
	_, err := ctx.Session.ChannelMessageSend(ctx.ChannelID, OutputString)
	return errors.Wrap(err, "couldn't reply")
}
func Tech(ctx *command.Context) error {
	data := [][]string{
		{"0", "None", "0", "1"},
		{"1", "Spear", "300", "1.05"},
		{"2", "Fire", "350", "1.10"},
		{"3", "Oven", "400", "1.16"},
		{"4", "Pottery", "460", "1.22"},
		{"5", "Domestication", "520", "1.28"},
		{"6", "Copper", "600", "1.34"},
		{"7", "Wheel", "690", "1.41"},
		{"8", "Writing", "800", "1.48"},
		{"9", "Bronze", "920", "1.55"},
		{"10", "Irrigation", "1,060", "1.63"},
		{"11", "Woodworking", "1,210", "1.71"},
		{"12", "Archery", "1,400", "1.80"},
		{"13", "Salt", "1,610", "1.89"},
		{"14", "Sailing", "1,850", "1.98"},
		{"15", "Masonary", "2,120", "2.08"},
		{"16", "Forum", "2,440", "2.18"},
		{"17", "Furnace", "2,810", "2.29"},
		{"18", "Ironworking", "3,230", "2.41"},
		{"19", "Library", "3,710", "2.53"},
		{"20", "Medicine", "4,270", "2.65"},
	}
	data2 := [][]string{
		{"21", "Timekeeping", "4,910", "2.79"},
		{"22", "Market", "5,650", "2.93"},
		{"23", "Monastary", "6,490", "3.07"},
		{"24", "Windmill", "7,470", "3.23"},
		{"25", "Printing", "8,590", "3.39"},
		{"26", "Civil Code", "9,880", "3.56"},
		{"27", "Shipbuilding", "11,360", "3.73"},
		{"28", "Astronomy", "13,060", "3.92"},
		{"29", "Chemistry", "15,020", "4.12"},
		{"30", "Gunpowder", "17,270", "4.32"},
		{"31", "Economics", "19,860", "4.54"},
		{"32", "Cotton Gin", "22,840", "4.78"},
		{"33", "Ballistics", "26,270", "5.00"},
		{"34", "Metallurgy", "30,210", "5.25"},
		{"35", "Laboratory", "34,740", "5.52"},
		{"36", "Mechanics", "39,950", "5.79"},
		{"37", "Textiles", "45,950", "6.08"},
		{"38", "Thermodynamics", "52,840", "6.39"},
		{"39", "Steam Engine", "60,760", "6.70"},
		{"40", "Assembly Line", "69,880", "7.04"},
		{"41", "Electricity", "80,360", "7.39"},
	}
	output := &bytes.Buffer{}
	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"Level", "Technology Name", "Cost (in EXP)", "Multiplier"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	output1 := &bytes.Buffer{}
	table1 := tablewriter.NewWriter(output1)
	table.SetHeader([]string{"Level", "Technology Name", "Cost (in EXP)", "Multiplier"})

	for _, v := range data2 {
		table1.Append(v)
	}
	table1.Render()

	OutputString := "```\n" + output.String() + "```"
	_, err := ctx.Session.ChannelMessageSend(ctx.ChannelID, OutputString)
	if err != nil {
		return errors.Wrap(err, "couldn't reply")
	}
	OutputString1 := "```\n" + output1.String() + "```"
	_, err = ctx.Session.ChannelMessageSend(ctx.ChannelID, OutputString1)
	return errors.Wrap(err, "couldn't reply")
}

func Econ(ctx *command.Context) error {
	data := [][]string{
		{"1", "Hunting", "300", "300"},
		{"2", "Farming", "600", "750"},
		{"3", "Fishing", "1,200", "1,950"},
		{"4", "Mining", "2,400", "4,650"},
		{"5", "Construction", "4,800", "11,700"},
		{"6", "Feudal", "9,600", "26,250"},
		{"7", "Trade", "19,200", "73,200"},
		{"8", "Exploration", "38,400", "183,150"},
		{"9", "Imperial", "76,800", "457,800"},
		{"10", "Mercantile", "153,600", "1,144,350"},
		{"11", "Plantation", "307,200", "2,861,100"},
		{"12", "Industrial", "614,400", "7,152,600"},
	}
	output := &bytes.Buffer{}
	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"Level", "Economic Development", "EXP", "Gold Per Minute"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	OutputString := "```\n" + output.String() + "```"
	_, err := ctx.Session.ChannelMessageSend(ctx.ChannelID, OutputString)
	return errors.Wrap(err, "couldn't reply")
}

func Exp(ctx *command.Context) error {
	data := [][]string{
		{"1", "1 EXP Per Min", "50,000,000"},
		{"2", "2 EXP Per Min", "150,000,000"},
		{"3", "3 EXP Per Min", "400,000,000"},
		{"4", "4 EXP Per Min", "600,000,000"},
		{"5", "5 EXP Per Min", "800,000,000"},
	}
	output := &bytes.Buffer{}
	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"Level", "Multiplier", "Cost (in Gold)"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	OutputString := "```\n" + output.String() + "```"
	_, err := ctx.Session.ChannelMessageSend(ctx.ChannelID, OutputString)
	return errors.Wrap(err, "couldn't reply")
}

func Safe(ctx *command.Context) error {
	data := [][]string{
		{"1", "5% Extra Per TBG/Min", "5,000"},
		{"2", "15% Extra Per TBG/Min", "7,500"},
		{"3", "20% Extra Per TBG/Min", "17,500"},
		{"4", "25% Extra Per TBG/Min", "25,000"},
		{"5", "30% Extra Per TBG/Min", "40,000"},
	}
	output := &bytes.Buffer{}
	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"Level", "Percentage", "Cost (in EXP)"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	OutputString := "```\n" + output.String() + "```"
	_, err := ctx.Session.ChannelMessageSend(ctx.ChannelID, OutputString)
	return errors.Wrap(err, "couldn't reply")
}
