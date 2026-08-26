package content

import "miaoxiu.example/domain"

func Seed() domain.Catalog {
	return domain.Catalog{Patterns: []domain.Pattern{
		{"p1", "蝴蝶妈妈", "生命起源与繁衍", "黔东南", "/static/butterfly.jpg", true}, {"p2", "龙纹", "守护与力量", "雷山", "/static/dragon.jpg", true}, {"p3", "鸟纹", "自由与迁徙", "台江", "/static/bird.jpg", false}, {"p4", "水波纹", "丰收与灵动", "丹寨", "/static/water.jpg", false},
	}, Stitches: []domain.Stitch{{"s1", "平绣", "针脚平整密集", "入门", "以短针并列填色"}, {"s2", "锁绣", "链状连续线迹", "进阶", "线圈相扣形成轮廓"}, {"s3", "辫绣", "编织感强", "进阶", "双线交错如发辫"}}, Artisans: []domain.Artisan{{"a1", "杨阿婆", "控拜村", "从十岁起学习苗绣，擅长蝴蝶纹。", "/static/yang.jpg", true}, {"a2", "石引梅", "新桥村", "国家级传承人，致力于年轻人教学。", "/static/shi.jpg", true}}, Artworks: []domain.Artwork{{"w1", "蝶舞银装", "p1", "a1", "/static/work1.jpg", "节庆盛装上的蝴蝶妈妈", 2022}, {"w2", "山河守望", "p2", "a2", "/static/work2.jpg", "龙纹与山水结合的挂饰", 2023}}}
}
