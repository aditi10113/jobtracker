let jobs=[]; const $=s=>document.querySelector(s);
async function api(url,opts={}){const r=await fetch(url,opts);const d=await r.json();if(!r.ok)throw Error(d.error||"Request failed");return d}
async function refresh(){jobs=await api("/api/jobs");render();const s=await api("/api/stats");$("#stats").innerHTML=`
<div class="stat"><strong>${s.total}</strong><span>Total applications</span></div>
<div class="stat"><strong>${s.applied}</strong><span>Applied</span></div>
<div class="stat"><strong>${s.screening}</strong><span>Screening</span></div>
<div class="stat"><strong>${s.interview}</strong><span>Interviews</span></div>
<div class="stat"><strong>${s.offer}</strong><span>Offers</span></div>
<div class="stat rate"><strong>${s.successRate.toFixed(0)}%</strong><span>Interview rate</span></div>`}
function render(){const q=$("#search").value.toLowerCase(),st=$("#statusFilter").value,pr=$("#priorityFilter").value;
const list=jobs.filter(j=>(!q||`${j.company} ${j.role} ${j.location}`.toLowerCase().includes(q))&&(!st||j.status===st)&&(!pr||j.priority===pr));
$("#jobs").innerHTML=list.length?list.map(j=>`<article class="job">
<div><div class="company">${esc(j.company)}</div><div class="role">${esc(j.role)}</div></div>
<div class="meta">📍 ${esc(j.location||"Not specified")}<br>${j.appliedDate?"Applied "+esc(j.appliedDate):""}</div>
<div><span class="status ${j.status.toLowerCase()}">${esc(j.status)}</span><div class="priority ${j.priority.toLowerCase()}">● ${esc(j.priority)} priority</div></div>
<div class="job-actions"><button class="icon" onclick="editJob(${j.id})">✎</button><button class="icon" onclick="deleteJob(${j.id})">×</button></div>
</article>`).join(""):`<div class="empty">No applications match your filters.</div>`}
function esc(v){return String(v??"").replace(/[&<>"']/g,c=>({"&":"&amp;","<":"&lt;",">":"&gt;",'"':"&quot;","'":"&#039;"}[c]))}
function openModal(j=null){$("#modal").classList.remove("hidden");$("#modalTitle").textContent=j?"Edit application":"Add job";const f=$("#jobForm");f.reset();f.id.value=j?.id||"";if(j)Object.keys(j).forEach(k=>{if(f.elements[k])f.elements[k].value=j[k]??""})}
function closeModal(){$("#modal").classList.add("hidden");$("#formError").textContent=""}
async function editJob(id){openModal(jobs.find(j=>j.id===id))}
async function deleteJob(id){if(!confirm("Delete this application?"))return;try{await api("/api/jobs/"+id,{method:"DELETE"});refresh()}catch(e){alert(e.message)}}
$("#openAdd").onclick=()=>openModal();$("#closeModal").onclick=closeModal;
$("#modal").addEventListener("click",e=>{if(e.target.id==="modal")closeModal()});
["search","statusFilter","priorityFilter"].forEach(id=>$("#"+id).addEventListener("input",render));
$("#jobForm").addEventListener("submit",async e=>{e.preventDefault();const f=new FormData(e.target);const data=Object.fromEntries(f.entries());delete data.id;const id=e.target.id.value;
try{await api(id?"/api/jobs/"+id:"/api/jobs",{method:id?"PUT":"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify(data)});closeModal();await refresh()}catch(err){$("#formError").textContent=err.message}});
refresh().catch(e=>{$("#jobs").innerHTML=`<div class="empty">Could not connect to the Go backend. Run <b>go run ./cmd/server</b>.</div>`});
