When any of the websites does not respond, all the goroutines should wait for the one until its timeout.

```console
$ go run main.go 
0.21s   24549 http://smatch.jp
0.29s   30573 https://www.recruit.co.jp
0.30s   33759 https://www.rikunabi.com
0.35s   11983 https://job.rikunabi.com
0.38s  232252 https://hikkoshi.suumo.jp
0.42s    4637 https://haken.rikunabi.com
0.44s   20219 https://connect.airregi.jp
0.44s   41117 https://townwork.net
0.45s  280560 https://www.recruit-dc.co.jp
0.46s    7191 https://toranet.jp
0.46s   40199 https://zexybaby.zexy.net
0.48s  118900 https://www.froma.com
0.48s   75321 https://airregi.jp
0.48s    4310 https://www.hatalike.jp
0.50s   24463 https://www.isize.com
0.53s   90414 https://recruit-card.jp
0.56s   70109 https://dock.cocokarada.jp
0.57s   39747 https://ats.joboplite.jp
0.60s   76711 https://twitter.com
0.60s   86204 https://www.hotpepper.jp
0.63s   35755 https://studysapuri.jp
0.65s   75420 https://www.suumocounter.jp
0.66s  122036 https://tabroom.jp
0.67s   75321 https://cdn.airregi.jp
0.67s   66727 https://rikunabi-yakuzaishi.jp
0.72s  121773 https://market.airregi.jp
0.72s   85564 https://suumo.jp
0.72s  226855 https://www.r-agent.com
0.72s  117347 https://beauty.hotpepper.jp
0.73s  131864 https://kaitori.carsensor.net
0.73s  188449 https://www.carsensor.net
0.75s  204003 https://www.jalan.net
0.77s   57651 https://faq.airregi.jp
0.78s  449744 https://www.ponparemall.com
0.81s  142279 https://www.r-staffing.co.jp
0.83s   13518 https://next.rikunabi.com
0.87s  193715 https://zexy.net
0.93s      14 https://app.adjust.com
0.98s  562796 https://youtu.be
1.01s  225712 https://www.facebook.com
1.07s  105627 https://shingakunet.com
1.12s 1298086 https://www.rgf-hragent.asia
1.31s  112317 https://comdeclab.com
1.95s  194146 https://golf-jalan.net
2.48s   90910 https://kaiten-portal.jp
2.478040s elapsed
```