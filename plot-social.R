# Ctrl + Alt + R ejecuta todo el script

setwd("~/go/src/github.com/estv-admin/find/")
# getwd()

df = read.csv("social.csv")
# df
# summary(df)
# str(df)
# plot(df)

library(ggplot2)

ggplot(df, aes(x="", y=value, fill=class)) +
  geom_bar(stat="identity", width=1, color="white") + 
  coord_polar("y", start=1) + 
  theme_void() # remove background, grid, numeric labels
