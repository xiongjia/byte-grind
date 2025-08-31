#ifndef _COMPONENT_HXX_
#define _COMPONENT_HXX_ 1

#include <QWidget>
#include <QPushButton>

class MyWindow : public QWidget
{
    Q_OBJECT
public:
    explicit MyWindow(QWidget *parent = 0);
private:
    QPushButton *m_button;
};

#endif // _COMPONENT_HXX_
