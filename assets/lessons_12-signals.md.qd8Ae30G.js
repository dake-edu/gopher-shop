import{_ as a,o as n,c as i,ae as e}from"./chunks/framework.CfVLvsAT.js";const c=JSON.parse('{"title":"Visual Signals Reference 🚦","description":"","frontmatter":{},"headers":[],"relativePath":"lessons/12-signals.md","filePath":"lessons/12-signals.md"}'),t={name:"lessons/12-signals.md"};function l(p,s,h,r,o,E){return n(),i("div",null,[...s[0]||(s[0]=[e(`<h1 id="visual-signals-reference-🚦" tabindex="-1">Visual Signals Reference 🚦 <a class="header-anchor" href="#visual-signals-reference-🚦" aria-label="Permalink to &quot;Visual Signals Reference 🚦&quot;">​</a></h1><p>This page uses the <strong>Shatalov Method</strong> of &quot;Visual Signals&quot; to explain complex Go concepts. Memorize these images, not the text.</p><h2 id="_1-the-interface-the-power-socket" tabindex="-1">1. The Interface (The Power Socket) <a class="header-anchor" href="#_1-the-interface-the-power-socket" aria-label="Permalink to &quot;1. The Interface (The Power Socket)&quot;">​</a></h2><p><strong>Concept</strong>: Decoupling implementation from usage. <strong>Signal</strong>: A Universal Power Socket. The Lamp doesn&#39;t care if the power comes from a wall, a battery, or a generator, as long as the plug fits.</p><div class="language-mermaid vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">mermaid</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">graph LR</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    subgraph Client [The User]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Lamp[&quot;🔌 Lamp (Needs Power)&quot;]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    end</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    subgraph Interface [The Interface]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Socket((&quot;Socket (Method: GivePower)&quot;))</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    end</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    subgraph Implementations [Concrete Types]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Wall[&quot;Wall Outlet&quot;]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Battery[&quot;Battery Pack&quot;]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Generator[&quot;Diesel Generator&quot;]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    end</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    Lamp --&gt; Socket</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    Wall -.-&gt; Socket</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    Battery -.-&gt; Socket</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    Generator -.-&gt; Socket</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    style Socket fill:#f9f,stroke:#333</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    style Lamp fill:#ff9,stroke:#333</span></span></code></pre></div><hr><h2 id="_2-channels-the-conveyor-belt" tabindex="-1">2. Channels (The Conveyor Belt) <a class="header-anchor" href="#_2-channels-the-conveyor-belt" aria-label="Permalink to &quot;2. Channels (The Conveyor Belt)&quot;">​</a></h2><p><strong>Concept</strong>: Safe communication between concurrent Goroutines. <strong>Signal</strong>: A Factory Conveyor Belt. Workers (Goroutines) put items on, other workers take them off. No one fights over the item; the belt manages the handoff.</p><div class="language-mermaid vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">mermaid</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">sequenceDiagram</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    participant G1 as 👷 Goroutine 1 (Producer)</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    participant Ch as 📦 Channel (Conveyor)</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    participant G2 as 👷 Goroutine 2 (Consumer)</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    G1-&gt;&gt;Ch: Arrow (&quot;Make Item&quot;)</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    Note over Ch: Item travels safely...</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    Ch-&gt;&gt;G2: Arrow (&quot;Take Item&quot;)</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    G2-&gt;&gt;G2: Process Item</span></span></code></pre></div><hr><h2 id="_3-structs-vs-pointers-the-blueprint-vs-the-house" tabindex="-1">3. Structs vs Pointers (The Blueprint vs The House) <a class="header-anchor" href="#_3-structs-vs-pointers-the-blueprint-vs-the-house" aria-label="Permalink to &quot;3. Structs vs Pointers (The Blueprint vs The House)&quot;">​</a></h2><p><strong>Concept</strong>: Value types vs Reference types. <strong>Signal</strong>: A Photocopy vs The Shared Document.</p><div class="language-mermaid vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">mermaid</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">graph TD</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    subgraph Value_Type [Struct (Value)]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Original[📄 Document A]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Copy[📄 Document B (Copy)]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Original -- &quot;Copying&quot; --&gt; Copy</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Note[If I edit Copy, Original is UNCHANGED]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    end</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    subgraph Pointer_Type [Pointer (*Struct)]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Shared[📄 shared_doc.txt]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Ref1[Ptr 1]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Ref2[Ptr 2]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Ref1 -- &quot;Points to&quot; --&gt; Shared</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Ref2 -- &quot;Points to&quot; --&gt; Shared</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        Note2[If Ptr 1 edits Doc, Ptr 2 SEES IT]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    end</span></span></code></pre></div><hr><h2 id="_4-the-quality-gate-validation" tabindex="-1">4. The Quality Gate (Validation) <a class="header-anchor" href="#_4-the-quality-gate-validation" aria-label="Permalink to &quot;4. The Quality Gate (Validation)&quot;">​</a></h2><p><strong>Concept</strong>: Fail early, fail fast. <strong>Signal</strong>: Airport Security. You don&#39;t get on the plane (Database) if you have a knife (Invalid Data).</p><div class="language-mermaid vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">mermaid</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">flowchart LR</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    Input[User Input] --&gt; Gate{👮 Security Gate}</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    </span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    Gate -- &quot;Invalid (No Name)&quot; --&gt; Reject[❌ 400 Bad Request]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    Gate -- &quot;Invalid (Price &lt; 0)&quot; --&gt; Reject</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    </span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    Gate -- &quot;Valid&quot; --&gt; Logic[🧠 Business Logic]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    Logic --&gt; DB[(🗄️ Database)]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    </span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    style Gate fill:#ffcccc,stroke:#r00</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    style DB fill:#ccffcc,stroke:#0f0</span></span></code></pre></div>`,17)])])}const g=a(t,[["render",l]]);export{c as __pageData,g as default};
