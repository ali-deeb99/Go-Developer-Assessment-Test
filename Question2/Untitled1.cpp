#include<bits/stdc++.h>
using namespace std;
int main(){


string s;
    cin >> s;

    vector<int> cnt(26);
    for(auto i: s) {
        ++cnt[i - 'a'];
    }

    priority_queue<pair<int, char>>pq;
    for(int i = 0; i < 26; i++) {
        if(cnt[i] != 0) {
            pq.push(make_pair(cnt[i], char(i + 'a')));
        }
    }

    int n = s.size();
    if(pq.top().first > (n + 1) / 2) {
        cout << "";
    }
    else {
        string res = "";
        while(pq.size() > 1) {
            auto [cnt1, c1] = pq.top();
            pq.pop();
            auto [cnt2, c2] = pq.top();
            pq.pop();
            res += c1;
            res += c2;
            if(cnt1 > 1) {
                pq.push(make_pair(cnt1 - 1, c1));
            }
            if(cnt2 > 1) {
                pq.push(make_pair(cnt2 - 1, c2));
            }
        }
        if(!pq.empty()) {
            auto [cnt1, c1] = pq.top();
            if(cnt1 > 1) {
                res = "";
            }
            else {
                res += c1;
            }
        }
        cout << res;
    }


    /*
int n,k,sum=0;
cin>>n>>k;

vector<int>nums(n);

for(int i=0;i<n;++i)
    cin>>nums[i];


vector<int>ans(n-k+1);

deque<int>dq;


dq.push_back(nums[0]);

for(int i=1;i<k;++i){
    while(!dq.empty()&&nums[i]>dq.back()){
        dq.pop_back();
    }
    dq.push_back(nums[i]);
}

ans[0]=dq.front();
int j=1;

for(int i=k;i<n;++i){
if(int(dq.size())==k)
dq.pop_front();
while(!dq.empty()&&nums[i]>dq.back()){
    dq.pop_back();
}
dq.push_back(nums[i]);
ans[j++]=dq.front();
}
for(int i=0;i<n-k+1;++i)
    cout<<ans[i]<<"\n";
*/
return 0;

}
