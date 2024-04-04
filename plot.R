# Ctrl + Alt + R ejecuta todo el script

setwd("~/go/src/github.com/estv-admin/find/")
# getwd()

df = read.csv("plot.csv")
# df
# summary(df)
# str(df)
# plot(df)

library(ggplot2)

ggplot(df, aes(x, y, colour = class)) + geom_point()
