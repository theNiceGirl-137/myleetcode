<p>给定一个仅包含数字&nbsp;<code>0-9</code>&nbsp;的字符串 <code>num</code> 和一个目标值整数 <code>target</code> ，在 <code>num</code> 的数字之间添加 <strong>二元 </strong>运算符（不是一元）<code>+</code>、<code>-</code>&nbsp;或&nbsp;<code>*</code>&nbsp;，返回所有能够得到目标值的表达式。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> <code>num = </code>"123", target = 6
<strong>输出: </strong>["1+2+3", "1*2*3"] 
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre>
<strong>输入:</strong> <code>num = </code>"232", target = 8
<strong>输出: </strong>["2*3+2", "2+3*2"]</pre>

<p><strong>示例 3:</strong></p>

<pre>
<strong>输入:</strong> <code>num = </code>"105", target = 5
<strong>输出: </strong>["1*0+5","10-5"]</pre>

<p><strong>示例&nbsp;4:</strong></p>

<pre>
<strong>输入:</strong> <code>num = </code>"00", target = 0
<strong>输出: </strong>["0+0", "0-0", "0*0"]
</pre>

<p><strong>示例 5:</strong></p>

<pre>
<strong>输入:</strong> <code>num = </code>"3456237490", target = 9191
<strong>输出: </strong>[]</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= num.length &lt;= 10</code></li>
	<li><code>num</code> 仅含数字</li>
	<li><code>-2<sup>31</sup> &lt;= target &lt;= 2<sup>31</sup> - 1</code></li>
</ul>
<div><div>Related Topics</div><div><li>数学</li><li>字符串</li><li>回溯</li></div></div><br><div><li>👍 284</li><li>👎 0</li></div>