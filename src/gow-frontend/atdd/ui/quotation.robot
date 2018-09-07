*** Settings ***
Library    SeleniumLibrary
Test Teardown    Close Browser

*** Variables ***
${HOST}    localhost:3000
${URL}    ${HOST}/quotation/create.html
${BROWSER}    chrome

*** Test Cases ***
ออกใบเสนอราคาโครงการ 3 Days Software Testing in Action Workshop
    เปิดหน้าเว็บสร้างใบเสนอราคา
    เลือกบริษัทที่ใช้ออกใบเสนอราคา
    กรอกข้อมูลบริษัทลูกค้า
    กรอกข้อมูลผู้ติดต่อฝั่งลูกค้า
    กรอกชื่อโครงการของการเสนอราคา
    กรอกข้อมูลรายละเอียดของโครงการ
    กรอกส่วนลด
    เลือกการวิธีการคิดภาษีมูลค่า
    กดปุ่มบันทึก
    แสดงข้อมูลรายละเอียดของใบเสนอราคา

*** Keywords ***
เปิดหน้าเว็บสร้างใบเสนอราคา
    Open Browser    ${URL}    ${BROWSER}
เลือกบริษัทที่ใช้ออกใบเสนอราคา
    Click Element    id=company_name
    Click Element    name=shuhari
กรอกข้อมูลบริษัทลูกค้า
    Input Text    id=customer_name    บริษัท อยุธยา แคปปิตอล เซอร์วิสเซส จำกัด
    Input Text    id=customer_branch    สำนักงานใหญ่
    Input Text    id=customer_address    อาคารกรุงศรีเพลินจิต ทาวเวอร์ 550 ถนนเพลินจิต แขวงเขตปทุมวัน กรุงเทพมหานคร
    Input Text    id=tax_id    0105537133562
กรอกข้อมูลผู้ติดต่อฝั่งลูกค้า 
    Input Text     id=name    Nopparat Slisatkorn
    Input Text    id=email    nopparat.slisatkorn@krungsri.com
กรอกชื่อโครงการของการเสนอราคา
    Input Text     id=project    โครงการ 3 Days Software Testing in Action Workshop
กรอกข้อมูลรายละเอียดของโครงการ
    Input Text     id=description    ค่าฝึกอบรม Software Testing in Action จำนวน 3 วัน วันพุธที่ 20 - วันศุกร์ที่ 22 มิถุนายน พ.ศ.2561
    Input Text    id=amount    1
    Input Text    id=priceperunit    100,000.00
    Input Text     id=price    100,000.00
กรอกส่วนลด
    Input Text    id=discount    0.00
เลือกการวิธีการคิดภาษีมูลค่า
    Click Element    id=outcome
กดปุ่มบันทึก
    Click Element    id=submit
แสดงข้อมูลรายละเอียดของใบเสนอราคา
    Wait Until Element Contains    id=company_name    บริษัท ชูฮาริ จำกัด
    Wait Until Element Contains    id=customer_name    บริษัท อยุธยา แคปปิตอล เซอร์วิสเซส จำกัด
    Wait Until Element Contains    id=customer_address    อาคารกรุงศรีเพลินจิต ทาวเวอร์ 550 ถนนเพลินจิต แขวงเขตปทุมวัน กรุงเทพมหานคร
    Wait Until Element Contains    id=tax_id    0105537133562
    Wait Until Element Contains    id=name    Nopparat Slisatkorn
    Wait Until Element Contains    id=email    nopparat.slisatkorn@krungsri.com
    Wait Until Element Contains    id=project    โครงการ 3 Days Software Testing in Action Workshop
    Wait Until Element Contains    id=description    ค่าฝึกอบรม Software Testing in Action จำนวน 3 วัน วันพุธที่ 20 - วันศุกร์ที่ 22 มิถุนายน พ.ศ.2561
    Wait Until Element Contains    id=amount    1
    Wait Until Element Contains    id=priceperunit    100,000.00
    Wait Until Element Contains    id=price    100,000.00
    Wait Until Element Contains    id=totalprice    100,000.00
    Wait Until Element Contains    id=discount    0.00
    Wait Until Element Contains    id=priceafterdiscount    100,000.00
    Wait Until Element Contains    id=tax_price    7,000.00
    Wait Until Element Contains    id=nettotal_price    107,000.00
    Wait Until Element Contains    id=thai_price    หนึ่งแสนเจ็ดพันบาทถ้วน

