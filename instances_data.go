// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2020-12-03 17:42:12.186689858 &#43;0000 UTC m=&#43;0.311114607
// using data from
// https://www.ec2instances.info/instances.json
package ec2instances

// Instances is a map from instance type to InstanceSize
var Instances = map[string]InstanceSize{
	"c5d.xlarge":    {4, 8.00},
	"m5dn.xlarge":   {4, 16.00},
	"c5.metal":      {96, 192.00},
	"d3en.8xlarge":  {32, 128.00},
	"m5a.2xlarge":   {8, 32.00},
	"r5n.12xlarge":  {48, 384.00},
	"c5.9xlarge":    {36, 72.00},
	"r5ad.xlarge":   {4, 32.00},
	"m5.24xlarge":   {96, 384.00},
	"r5n.xlarge":    {4, 32.00},
	"i3en.12xlarge": {48, 384.00},
	"i3en.metal":    {96, 768.00},
	"m5zn.3xlarge":  {12, 48.00},
	"m5d.12xlarge":  {48, 192.00},
	"c5.large":      {2, 4.00},
	"r5dn.xlarge":   {4, 32.00},
	"c5n.large":     {2, 5.25},
	"c5.24xlarge":   {96, 192.00},
	"i2.xlarge":     {4, 30.50},
	"m5n.16xlarge":  {64, 256.00},
	"t2.micro":      {1, 1.00},
	"d2.8xlarge":    {36, 244.00},
	"i3en.3xlarge":  {12, 96.00},
	"z1d.3xlarge":   {12, 96.00},
	"m5.2xlarge":    {8, 32.00},
	"inf1.xlarge":   {4, 8.00},
	"c5.18xlarge":   {72, 144.00},
	"r6gd.16xlarge": {64, 512.00},
	"x1e.16xlarge":  {64, 1952.00},
	"r5n.24xlarge":  {96, 768.00},
	"i2.8xlarge":    {32, 244.00},
	"r5a.8xlarge":   {32, 256.00},
	"r6g.medium":    {1, 8.00},
	"r6g.12xlarge":  {48, 384.00},
	"a1.metal":      {16, 32.00},
	"t4g.micro":     {2, 1.00},
	"m5zn.metal":    {48, 192.00},
	"i2.2xlarge":    {8, 61.00},
	"i3en.2xlarge":  {8, 64.00},
	"m5a.xlarge":    {4, 16.00},
	"p3.2xlarge":    {8, 61.00},
	"c6g.2xlarge":   {8, 16.00},
	"t2.2xlarge":    {8, 32.00},
	"h1.8xlarge":    {32, 128.00},
	"r5d.24xlarge":  {96, 768.00},
	"i3en.6xlarge":  {24, 192.00},
	"r4.8xlarge":    {32, 244.00},
	"t2.large":      {2, 8.00},
	"x1.16xlarge":   {64, 976.00},
	"c5d.18xlarge":  {72, 144.00},
	"m5a.16xlarge":  {64, 256.00},
	"r5.metal":      {96, 768.00},
	"r5a.large":     {2, 16.00},
	"c3.large":      {2, 3.75},
	"c6gd.medium":   {1, 2.00},
	"inf1.24xlarge": {96, 192.00},
	"r5a.24xlarge":  {96, 768.00},
	"g3.16xlarge":   {64, 488.00},
	"a1.2xlarge":    {8, 16.00},
	"c4.xlarge":     {4, 7.50},
	"x1e.4xlarge":   {16, 488.00},
	"m5ad.xlarge":   {4, 16.00},
	"i3en.24xlarge": {96, 768.00},
	"c5n.18xlarge":  {72, 192.00},
	"c6gd.metal":    {64, 128.00},
	"m4.large":      {2, 8.00},
	"t3.medium":     {2, 4.00},
	"h1.4xlarge":    {16, 64.00},
	"x1e.xlarge":    {4, 122.00},
	"g4dn.8xlarge":  {32, 128.00},
	"r5n.large":     {2, 16.00},
	"m5.large":      {2, 8.00},
	"c5.4xlarge":    {16, 32.00},
	"m5d.24xlarge":  {96, 384.00},
	"r3.large":      {2, 15.25},
	"r5n.2xlarge":   {8, 64.00},
	"c4.large":      {2, 3.75},
	"r5d.xlarge":    {4, 32.00},
	"m5d.large":     {2, 8.00},
	"r5.2xlarge":    {8, 64.00},
	"m3.2xlarge":    {8, 30.00},
	"m5n.12xlarge":  {48, 192.00},
	"m5d.2xlarge":   {8, 32.00},
	"m4.10xlarge":   {40, 160.00},
	"m5ad.large":    {2, 8.00},
	"c5d.large":     {2, 4.00},
	"r5dn.8xlarge":  {32, 256.00},
	"m6g.2xlarge":   {8, 32.00},
	"c5d.24xlarge":  {96, 192.00},
	"r5n.16xlarge":  {64, 512.00},
	"r5d.metal":     {96, 768.00},
	"x1.32xlarge":   {128, 1952.00},
	"c5n.2xlarge":   {8, 21.00},
	"r5n.8xlarge":   {32, 256.00},
	"d3en.xlarge":   {4, 16.00},
	"m5d.xlarge":    {4, 16.00},
	"m5.metal":      {96, 384.00},
	"m5dn.large":    {2, 8.00},
	"m5.12xlarge":   {48, 192.00},
	"p3.16xlarge":   {64, 488.00},
	"r6gd.12xlarge": {48, 384.00},
	"r6g.2xlarge":   {8, 64.00},
	"m3.large":      {2, 7.50},
	"c5d.4xlarge":   {16, 32.00},
	"m5n.2xlarge":   {8, 32.00},
	"r5.24xlarge":   {96, 768.00},
	"m5dn.12xlarge": {48, 192.00},
	"c6gd.2xlarge":  {8, 16.00},
	"r5dn.16xlarge": {64, 512.00},
	"r4.large":      {2, 15.25},
	"r3.xlarge":     {4, 30.50},
	"x1e.32xlarge":  {128, 3904.00},
	"c5n.xlarge":    {4, 10.50},
	"m5a.4xlarge":   {16, 64.00},
	"r5ad.2xlarge":  {8, 64.00},
	"t2.xlarge":     {4, 16.00},
	"p3dn.24xlarge": {96, 768.00},
	"p2.16xlarge":   {64, 732.00},
	"m5dn.16xlarge": {64, 256.00},
	"c3.8xlarge":    {32, 60.00},
	"m3.medium":     {1, 3.75},
	"x1e.2xlarge":   {8, 244.00},
	"m5dn.24xlarge": {96, 384.00},
	"m5a.8xlarge":   {32, 128.00},
	"m5ad.2xlarge":  {8, 32.00},
	"m5zn.6xlarge":  {24, 96.00},
	"d3en.6xlarge":  {24, 96.00},
	"r5a.2xlarge":   {8, 64.00},
	"c5d.metal":     {96, 192.00},
	"m5a.12xlarge":  {48, 192.00},
	"t2.small":      {1, 2.00},
	"m5d.16xlarge":  {64, 256.00},
	"d2.xlarge":     {4, 30.50},
	"c5.2xlarge":    {8, 16.00},
	"g4dn.16xlarge": {64, 256.00},
	"m6g.medium":    {1, 4.00},
	"m1.xlarge":     {4, 15.00},
	"r5b.12xlarge":  {48, 384.00},
	"m5n.xlarge":    {4, 16.00},
	"m6g.16xlarge":  {64, 256.00},
	"c5d.2xlarge":   {8, 16.00},
	"m3.xlarge":     {4, 15.00},
	"m1.medium":     {1, 3.75},
	"t3a.medium":    {2, 4.00},
	"r5b.large":     {2, 16.00},
	"r4.16xlarge":   {64, 488.00},
	"i3.xlarge":     {4, 30.50},
	"z1d.2xlarge":   {8, 64.00},
	"c3.xlarge":     {4, 7.50},
	"r5.8xlarge":    {32, 256.00},
	"c3.2xlarge":    {8, 15.00},
	"r3.2xlarge":    {8, 61.00},
	"r4.2xlarge":    {8, 61.00},
	"m6gd.12xlarge": {48, 192.00},
	"r6g.8xlarge":   {32, 256.00},
	"d3.8xlarge":    {32, 256.00},
	"p2.8xlarge":    {32, 488.00},
	"m5.xlarge":     {4, 16.00},
	"c4.8xlarge":    {36, 60.00},
	"c5n.9xlarge":   {36, 96.00},
	"c5a.8xlarge":   {32, 64.00},
	"d2.2xlarge":    {8, 61.00},
	"r6g.metal":     {64, 512.00},
	"d2.4xlarge":    {16, 122.00},
	"m5ad.4xlarge":  {16, 64.00},
	"c5.xlarge":     {4, 8.00},
	"c6g.4xlarge":   {16, 32.00},
	"m5dn.4xlarge":  {16, 64.00},
	"inf1.6xlarge":  {24, 48.00},
	"g4dn.metal":    {96, 384.00},
	"d3en.12xlarge": {48, 192.00},
	"d3.4xlarge":    {16, 128.00},
	"d3.2xlarge":    {8, 64.00},
	"r6gd.2xlarge":  {8, 64.00},
	"m5zn.12xlarge": {48, 192.00},
	"c5a.xlarge":    {4, 8.00},
	"m5d.4xlarge":   {16, 64.00},
	"r5.16xlarge":   {64, 512.00},
	"m5n.24xlarge":  {96, 384.00},
	"m4.xlarge":     {4, 16.00},
	"r5a.4xlarge":   {16, 128.00},
	"t3.xlarge":     {4, 16.00},
	"z1d.12xlarge":  {48, 384.00},
	"r5b.metal":     {96, 768.00},
	"m4.16xlarge":   {64, 256.00},
	"r5.12xlarge":   {48, 384.00},
	"m4.4xlarge":    {16, 64.00},
	"r4.xlarge":     {4, 30.50},
	"m1.small":      {1, 1.70},
	"c5ad.8xlarge":  {32, 64.00},
	"a1.xlarge":     {4, 8.00},
	"z1d.6xlarge":   {24, 192.00},
	"r5b.8xlarge":   {32, 256.00},
	"r5d.4xlarge":   {16, 128.00},
	"r6gd.8xlarge":  {32, 256.00},
	"c6g.16xlarge":  {64, 128.00},
	"r6g.16xlarge":  {64, 512.00},
	"m5dn.metal":    {96, 384.00},
	"r5a.16xlarge":  {64, 512.00},
	"m6gd.large":    {2, 8.00},
	"t4g.large":     {2, 8.00},
	"m5zn.2xlarge":  {8, 32.00},
	"m5n.large":     {2, 8.00},
	"m5dn.8xlarge":  {32, 128.00},
	"p2.xlarge":     {4, 61.00},
	"c6gd.16xlarge": {64, 128.00},
	"c3.4xlarge":    {16, 30.00},
	"r4.4xlarge":    {16, 122.00},
	"r5a.xlarge":    {4, 32.00},
	"m5n.4xlarge":   {16, 64.00},
	"h1.2xlarge":    {8, 32.00},
	"m5ad.24xlarge": {96, 384.00},
	"c6g.xlarge":    {4, 8.00},
	"m6g.metal":     {64, 256.00},
	"d3en.4xlarge":  {16, 64.00},
	"f1.4xlarge":    {16, 244.00},
	"c5a.24xlarge":  {96, 192.00},
	"i3en.xlarge":   {4, 32.00},
	"c5a.4xlarge":   {16, 32.00},
	"r5a.12xlarge":  {48, 384.00},
	"m5zn.large":    {2, 8.00},
	"u-12tb1.metal": {448, 12288.00},
	"r5ad.24xlarge": {96, 768.00},
	"c5ad.4xlarge":  {16, 32.00},
	"r6gd.xlarge":   {4, 32.00},
	"m5d.metal":     {96, 384.00},
	"g2.2xlarge":    {8, 15.00},
	"c4.2xlarge":    {8, 15.00},
	"x1e.8xlarge":   {32, 976.00},
	"r5dn.large":    {2, 16.00},
	"m5.4xlarge":    {16, 64.00},
	"m6gd.16xlarge": {64, 256.00},
	"r5dn.12xlarge": {48, 384.00},
	"r5n.4xlarge":   {16, 128.00},
	"h1.16xlarge":   {64, 256.00},
	"z1d.xlarge":    {4, 32.00},
	"r5dn.4xlarge":  {16, 128.00},
	"c5d.12xlarge":  {48, 96.00},
	"c5ad.24xlarge": {96, 192.00},
	"r5d.8xlarge":   {32, 256.00},
	"t3.2xlarge":    {8, 32.00},
	"r6gd.4xlarge":  {16, 128.00},
	"d3en.2xlarge":  {8, 32.00},
	"g3.8xlarge":    {32, 244.00},
	"m6gd.medium":   {1, 4.00},
	"m5d.8xlarge":   {32, 128.00},
	"m6gd.4xlarge":  {16, 64.00},
	"c6gd.large":    {2, 4.00},
	"m5zn.xlarge":   {4, 16.00},
	"i3.metal":      {64, 512.00},
	"m5.16xlarge":   {64, 256.00},
	"t4g.small":     {2, 2.00},
	"m4.2xlarge":    {8, 32.00},
	"c5a.2xlarge":   {8, 16.00},
	"z1d.metal":     {48, 384.00},
	"g4dn.12xlarge": {48, 192.00},
	"r5ad.large":    {2, 16.00},
	"r5b.4xlarge":   {16, 128.00},
	"r5.4xlarge":    {16, 128.00},
	"r5ad.8xlarge":  {32, 256.00},
	"c6g.metal":     {64, 128.00},
	"r5d.2xlarge":   {8, 64.00},
	"m5ad.8xlarge":  {32, 128.00},
	"r5.large":      {2, 16.00},
	"i3.large":      {2, 15.25},
	"z1d.large":     {2, 16.00},
	"m6g.4xlarge":   {16, 64.00},
	"m6gd.xlarge":   {4, 16.00},
	"i3.16xlarge":   {64, 488.00},
	"m5ad.12xlarge": {48, 192.00},
	"r5b.16xlarge":  {64, 512.00},
	"i3en.large":    {2, 16.00},
	"r5ad.16xlarge": {64, 512.00},
	"c5ad.large":    {2, 4.00},
	"c5n.metal":     {72, 192.00},
	"i2.4xlarge":    {16, 122.00},
	"m5ad.16xlarge": {64, 256.00},
	"m5dn.2xlarge":  {8, 32.00},
	"c5d.9xlarge":   {36, 72.00},
	"r5.xlarge":     {4, 32.00},
	"i3.2xlarge":    {8, 61.00},
	"g4dn.xlarge":   {4, 16.00},
	"c5a.12xlarge":  {48, 96.00},
	"r5d.large":     {2, 16.00},
	"g3.4xlarge":    {16, 122.00},
	"r5dn.24xlarge": {96, 768.00},
	"r5d.12xlarge":  {48, 384.00},
	"r5dn.2xlarge":  {8, 64.00},
	"g2.8xlarge":    {32, 60.00},
	"i3.4xlarge":    {16, 122.00},
	"g4dn.4xlarge":  {16, 64.00},
	"c5.12xlarge":   {48, 96.00},
	"c5n.4xlarge":   {16, 42.00},
	"t3a.small":     {2, 2.00},
	"r3.4xlarge":    {16, 122.00},
	"m1.large":      {2, 7.50},
	"r5b.2xlarge":   {8, 64.00},
	"c5ad.xlarge":   {4, 8.00},
	"m2.4xlarge":    {8, 68.40},
	"inf1.2xlarge":  {8, 16.00},
	"r3.8xlarge":    {32, 244.00},
	"p3.8xlarge":    {32, 244.00},
	"m6gd.metal":    {64, 256.00},
	"c5ad.2xlarge":  {8, 16.00},
	"m2.xlarge":     {2, 17.10},
	"c1.medium":     {2, 1.70},
	"r6gd.metal":    {64, 512.00},
	"m6g.8xlarge":   {32, 128.00},
	"cc2.8xlarge":   {32, 60.50},
	"c5ad.16xlarge": {64, 128.00},
	"t3a.2xlarge":   {8, 32.00},
	"c1.xlarge":     {8, 7.00},
	"c5ad.12xlarge": {48, 96.00},
	"c5a.large":     {2, 4.00},
	"c6g.medium":    {1, 2.00},
	"r5d.16xlarge":  {64, 512.00},
	"m5.8xlarge":    {32, 128.00},
	"m2.2xlarge":    {4, 34.20},
	"r5ad.4xlarge":  {16, 128.00},
	"c6gd.8xlarge":  {32, 64.00},
	"t3a.nano":      {2, 0.50},
	"m5n.metal":     {96, 384.00},
	"c6gd.12xlarge": {48, 96.00},
	"c5a.16xlarge":  {64, 128.00},
	"r5ad.12xlarge": {48, 384.00},
	"c6gd.4xlarge":  {16, 32.00},
	"r5b.24xlarge":  {96, 768.00},
	"m5a.large":     {2, 8.00},
	"i3.8xlarge":    {32, 244.00},
	"cr1.8xlarge":   {32, 244.00},
	"c6g.12xlarge":  {48, 96.00},
	"m5n.8xlarge":   {32, 128.00},
	"m6gd.8xlarge":  {32, 128.00},
	"m5a.24xlarge":  {96, 384.00},
	"t3.nano":       {2, 0.50},
	"d3.xlarge":     {4, 32.00},
	"g4dn.2xlarge":  {8, 32.00},
	"r6g.4xlarge":   {16, 128.00},
	"c4.4xlarge":    {16, 30.00},
	"t3.micro":      {2, 1.00},
	"u-6tb1.metal":  {448, 6144.00},
	"t3.small":      {2, 2.00},
	"m6gd.2xlarge":  {8, 32.00},
	"f1.16xlarge":   {64, 976.00},
	"t3.large":      {2, 8.00},
	"t3a.micro":     {2, 1.00},
	"u-9tb1.metal":  {448, 9216.00},
	"c6g.8xlarge":   {32, 64.00},
	"g3s.xlarge":    {4, 30.50},
	"p4d.24xlarge":  {96, 1152.00},
	"r5b.xlarge":    {4, 32.00},
	"r6g.large":     {2, 16.00},
	"a1.4xlarge":    {16, 32.00},
	"f1.2xlarge":    {8, 122.00},
	"a1.medium":     {1, 2.00},
	"t2.medium":     {2, 4.00},
	"t4g.2xlarge":   {8, 32.00},
	"a1.large":      {2, 4.00},
	"r6g.xlarge":    {4, 32.00},
	"hs1.8xlarge":   {16, 117.00},
	"u-18tb1.metal": {448, 18432.00},
	"t3a.xlarge":    {4, 16.00},
	"c6g.large":     {2, 4.00},
	"t1.micro":      {1, 0.61},
	"r6gd.large":    {2, 16.00},
	"m6g.xlarge":    {4, 16.00},
	"m6g.12xlarge":  {48, 192.00},
	"r5dn.metal":    {96, 768.00},
	"t4g.xlarge":    {4, 16.00},
	"m6g.large":     {2, 8.00},
	"r5n.metal":     {96, 768.00},
	"t4g.medium":    {2, 4.00},
	"c6gd.xlarge":   {4, 8.00},
	"t4g.nano":      {2, 0.50},
	"r6gd.medium":   {1, 8.00},
	"u-24tb1.metal": {448, 24576.00},
	"t2.nano":       {1, 0.50},
	"t3a.large":     {2, 8.00},
	"mac1.metal":    {12, 32.00},
}
